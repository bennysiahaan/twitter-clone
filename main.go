package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bennysiahaan/twitter-clone/data"
	"github.com/bennysiahaan/twitter-clone/db"
	"github.com/bennysiahaan/twitter-clone/handlers"
	"github.com/go-openapi/runtime/middleware"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var BindAddress = ":9090"

func main() {
    
    l := log.New(os.Stdout, "twitter-clone: ", log.LstdFlags)
	v := data.NewTweetValidation()
    
	TweetHandler := handlers.NewTweet(l, v)
    
    // CORS
    CORSHandler := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	router := mux.NewRouter()

	// API handlers
	GetHandler := router.Methods(http.MethodGet, http.MethodOptions).Subrouter()
	GetHandler.HandleFunc("/home", TweetHandler.GetTimeline)

	GetHandler.HandleFunc("/tweet/{tweetId:[0-9a-zA-Z-]{36}}", TweetHandler.GetTweet)
    
	PostHandler := router.Methods(http.MethodPost).Subrouter()
	PostHandler.HandleFunc("/create", TweetHandler.Post)
	PostHandler.Use(TweetHandler.MiddlewareValidateTweet)
    
	PutHandler := router.Methods(http.MethodPut).Subrouter()
	PutHandler.HandleFunc("/edit", TweetHandler.Edit)
	PutHandler.Use(TweetHandler.MiddlewareValidateTweet)
    
	DeleteHandler := router.Methods(http.MethodDelete).Subrouter()
	DeleteHandler.HandleFunc("/tweet/{tweetId:[0-9a-zA-Z-]{36}}", TweetHandler.Delete)
    
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)
    
	GetHandler.Handle("/docs", sh)
	GetHandler.Handle("/swagger.yaml", http.FileServer(http.Dir("./docs")))
    
    // Redirects
	GetHandler.HandleFunc("/{tweetId:[0-9a-zA-Z-]{36}}", TweetHandler.RedirectTweet)
	GetHandler.HandleFunc("/", TweetHandler.RedirectHome)
    
	srv := &http.Server{
		Addr:         BindAddress,
		Handler:      CORSHandler(router),
		ErrorLog:     l,
		ReadTimeout:  time.Duration(5) * time.Second,
		WriteTimeout: time.Duration(10) * time.Second,
		IdleTimeout:  time.Duration(120) * time.Second,
	}

	go func() {
		// connect to mysql db
		db.ConnectDB(l)
		DB := db.GetDB()
		defer DB.Close()
		if err := DB.Ping(); err != nil {
			l.Fatal(err)
		}

		l.Println("Starting server on port 9090")

		err := srv.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
            os.Exit(1)
		}
	}()

	// trap sigterm or interrupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	sig := <-c
	l.Println("Got signal:", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
	l.Println("Shutting down...")
	os.Exit(0)

}
