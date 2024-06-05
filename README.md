# go-rest-api

RESTful API using Golang with the Gin framework and GORM for database operations. JWT is used for authentication, and Swagger is integrated for API documentation.

## Table of Contents

1. [Database Setup](#database-setup)
2. [DB Migrations](#db-migrations)
3. [API Documentation](#api-documentation)
4. [Setup and Run the Application](#setup-and-run-the-application)
5. [API Endpoints](#api-endpoints)
6. [Unit Tests](#unit-tests)

## Database Setup

Run the following command to set up a PostgreSQL database using Docker:

```sh
docker run --name go-rest-api-db -e POSTGRES_USER=gigawrks -e POSTGRES_PASSWORD=dev2s -e POSTGRES_DB=user-db -p 5432:5432 -d postgres
```

### Database Credentials

-   Username: gigawrks
-   Password: dev2s
-   Database Name: user-db

## DB Migrations

Install migrate

Download and install the migrate tool:

```sh
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz | tar xvz

sudo mv migrate /usr/local/bin/migrate
```

## Connect to the Database

Use the following command to connect to the PostgreSQL database:

```sh
docker exec -it go-rest-api-db bash

psql postgresql://gigawrks:dev2s@localhost:5432/user-dbs
```

### Run migration scripts

create all the relations

```sh
./migrations/run-up.sh
```

rollback changes

```sh
./migrations/run-down.sh
```

## API Documentation

### Install Swagger

Install Swagger using the following command:

```sh
swag init --parseDependency --parseInternal -g main.go -o docs
```

### Generate Swagger Docs

Generate the Swagger documentation with:

```sh
swag init --parseDependency --parseInternal -g main.go -o docs
```

### Access Swagger UI

Once the application is running, you can access the Swagger UI at:

```sh
{base-url}/swagger/index.html
```

## Setup and Run the Application

### Prerequisites

    - Golang (version 1.16 or later)
    - Docker

### Clone the Repository

Clone this repository to your local machine:

```sh
git clone https://github.com/your-username/go-rest-api.git
cd go-rest-api
```

### Install Dependencies

Install the required Go dependencies:

```sh
go mod tidy
```

### Run the Application

Start the application using:

```sh
go run main.go
```

Ensure the database is up and running before starting the application, follow steps [here](#database-setup).

## Unit Tests

### Generating mocks

```sh
mockgen -source=services/user_service.go -destination=services/mocks/user_service_mock.go -package=services

mockgen -source=database/database.go -destination=database/mocks/database_mock.go -package=database
```
