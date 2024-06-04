# go-rest-api

RESTful API using Golang with the Gin framework and Gorm for database operations. Also integrate JWT for authentication and Swagger for API documentation.

# database setup

### username : gigawrks

### password : dev2s

### db name : user-db

docker run --name go-rest-api-db -e POSTGRES_USER=gigawrks -e POSTGRES_PASSWORD=dev2s -e POSTGRES_DB=user-db -p 5432:5432 -d postgres

# db migrations

### install `migrate`

curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz | tar xvz
sudo mv migrate.linux-amd64 /usr/local/bin/migrate

### connect

psql postgresql://gigawrks:dev2s@localhost:5432/user-db

# Swagger

go install github.com/swaggo/swag/cmd/swag@latest

swag init --parseDependency --parseInternal -g main.go -o docs

Endpoint : {base-url}/swagger/index.html
