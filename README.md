# Simple API Twitter Clone

A REST API microservices-based project to fetch, edit, post, and delete tweets.

## API documentation

The API documentation is built with go-swagger. The doc is accessed from `/docs` path in the URL.

## Build with Makefile

The build requires Go version [1.17.3](https://github.com/actions/go-versions/releases/tag/1.17.3-1425023214) or later to run and requires [go-swagger](https://github.com/go-swagger/go-swagger) to be installed in order to build the API documentation.

### Installing go-swagger

To install go-swagger, run:

```
make install_swagger
```

### Building API docs with go-swagger

To build the API docs, run:

```
make swagger
```

## SQL database schema

To use the default database schema, initialize your SQL database using `init.sql` under `mysql` directory. Run in your SQL server:

```sql
source $PATH\mysql\init.sql
```

and change `$PATH` with the full path to this project directory.
