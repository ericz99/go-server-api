# Go Server REST API Example

A RESTful API example for simple todo application with Go

## Installation & Run

```bash
# Download this project
go get github.com/ericz99/go-server-api
```

Before running API server, you should set the database config with yours or set the your database config with my values on [config.go](https://github.com/ericz99/go-server-api/blob/master/config/config.go)

Also, please change .env.example -> .env and update the all fields.

```bash
# production env mode
mode = development

# database config
db_name = testdb
db_pass = passwrod
db_user = root
db_type = mysql
db_host = localhost
db_port = 3306

# server port
port = 8080

# for jwt secret key
secret_key = feTATm1?@d+1GKG
```

```bash
# Build and Run
cd go-server-api
go build
./go-server-api - [if you\'re using linux/macos]
go-server-api - [if you\'re using windows]

# API Endpoint (VERSION 1) : http://localhost:8080/api/v1
# API Endpoint (VERSION 2) : http://localhost:8080/api/v2
# API Auth Endpoint : http://localhost:8080/api/auth
```

## API

## Auth Endpoint

#### /register

- `POST` : Create new user

#Post Params

```
{
	"name": "Hello World",
	"email": "test@yahoo.com",
	"password": "asdfasdf"
}
```

#### /login

- `POST` : Login User

#Post Params

```
{
	"email": "test@yahoo.com",
	"password": "asdfasdf"
}
```

## Protected Endpoint

Please add `x-auth-token` in your header in order to get authorized to any of these endpoint below!

#### /book/save

- `POST` : Save a book

#### /books

- `GET` : Get all book

#### /book/:id

- `GET` : Get a book
- `DELETE` : Delete a book

#Post Params

```
{
	"title": "LOL: Book PT 2",
	"isbn": "isbn-6s9",
	"author": {
		"name": "Bob"
	}
}
```

## Todo

- [x] Support basic REST APIs.
- [x] Support GORM/Database
- [x] Support Authentication with user for securing the APIs.
- [x] Make convenient wrappers for creating API handlers.
- [ ] Write the tests for all APIs.
- [x] Organize the code with packages
- [ ] Make docs with GoDoc
- [ ] Building a deployment process
