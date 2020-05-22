# Overview

### Todo

## Database

The project makes use of a PostgreSQL 12 database.

Migrations are handled via [golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

#### Creating a New Migration

```
migrate create -dir=db/migrations/ -ext=.sql <migration_name>
```

#### Running Migrations

```
migrate -path=db/migrations/ -database <postgres_connection_string> up
```

#### Running Tests

```
go test ./... -cover
```

