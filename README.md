# Go Server REST API Example

A RESTful API example for simple todo application with Go

## Installation & Run

```bash
# Download this project
go get github.com/ericz99/go-server-api
```

Before running API server, you should set the database config with yours or set the your database config with my values on [config.go](https://github.com/ericz99/go-server-api/blob/master/config/config.go)

```go
func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Username: "guest",
			Password: "Guest0000!",
			Name:     "todoapp",
			Charset:  "utf8",
		},
	}
}
```

```bash
# Build and Run
cd go-server-api
go build
./go-server-api - [if you\'re using linux/macos]
go-server-api - [if you\'re using windows]

# Please change .env.example -> .env and update the values

# API Endpoint (VERSION 1) : http://localhost:8080/api/v1
# API Endpoint (VERSION 2) : http://localhost:8080/api/v2
```

## API

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
