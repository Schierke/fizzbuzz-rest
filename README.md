# fizzbuzz-backend

## Introduction
Providing a fizzbuzz calculation using REST API

## Getting started
Before you start, make sure you have ```golang``` and ```docker``` installed with lastest version

For local execution, I recommend using [task](https://github.com/go-task/task).

## INSTALL

### Configure the app.env file

Start by configuring the app.env file to connect to your database and also the port application will be using. 

### Local development

```docker compose up -d``` : initialzing the DB. The database will be hosted on port 5432

```task run```: running complied project

### Database migration

The migration is done by using ```golang-migrate```. For each fresh use on local or deployed box, please run ```task migrate```

## Test

- You can generate new mock files by using ```task mock```

- Run the unit test by ```task test```

- You can also import the postman file from docs/postman for testing the entire application

### DEPENDENCIES

- chi - HTTP Services
- zerolog - logger
- Testify - tool for code testing
- pgx - create connection pool to PostgreSQL
- cobra - CLI
- gomock & pgxmock: mocks
