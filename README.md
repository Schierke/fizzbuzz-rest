# fizzbuzz-backend

## Introduction
Providing a fizzbuzz calculation using REST API

## Getting started
Before you start, make sure you have ```golang``` and ```docker``` installed with lastest version

## INSTALL

```docker compose up -d``` : initialzing the DB. The database will be hosted on port 5432

```go build``` : building service

```./fizzbuzz hydrate``` : database migration. It will destroy all the current DBs then create new one (testing and deploying purpose)

```./fizzbuzz serve``` : running the service. App listened on port 8080




### DEPENDENCIES

- chi - HTTP Services
- zerolog - logger
- Testify - tool for code testing
- pgx - create connection pool to PostgreSQL
- cobra - CLI
- gomock & pgxmock: mocks
