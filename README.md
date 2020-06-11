# Overview

An application which periodically reads and stores gauge data for rivers via the United States Geological Survey [Instantaneous Values Web Service](https://waterservices.usgs.gov/rest/IV-Service.html).  

The gauge data is served via a read only REST API. 

The API is intended to be consumed by a front end application aimed at providing easy to read river information for paddlers. 

## Endpoints

TODO Add Metrics Endpoints

`/api/rivers` - Returns a list of all tracked rivers

`/api/rivers/:id` - Returns a single river

`/api/rivers/:id/gauges` - Returns a list of gauges for a single river

`/api/gauges` - Returns a list of all tracked gauges

`/api/gauges/:id` - Returns a single gauge

## Running a Local Instance

A local instance can be created in docker using the docker compose in the root directory of the repo.

This will spin up an instance running the latest master.

```
docker-compose up
```

### Seeding the Database

The database can be seeded by running the `seeder` program in the `ort_api` container passing the directory of the seed files as a parameter.

```
docker exec -it ort_api /bin/bash

./seeder rivers/
```

### Reading Gauges

By default gauges are read, and their values stored every 15 minutes. TODO make this configurable with an env var.

## Database

The project makes use of a PostgreSQL 12 database.

Migrations are handled via [golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

#### Creating a New Migration

```
migrate create -dir=db/migrations/ -ext=.sql <migration_name>
```

#### Running Migrations

Migrations run automatically when starting the main Go API.

However, they can also be run manually if need be.

```
migrate -path=db/migrations/ -database <postgres_connection_string> up
```

## Tests

Tests are written using [test suites](https://pkg.go.dev/github.com/stretchr/testify/suite?tab=doc)

#### Running Tests

```
go test ./... -cover
```

