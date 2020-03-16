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
cd go-server-api-example
go build
./go-server-api-example - [if you\'re using linux/macos]
go-server-api-example - [if you\'re using windows]

# API Endpoint : http://localhost:8080
```

## Todo

- [x] Support basic REST APIs.
- [x] Support GORM/Database
- [ ] Support Authentication with user for securing the APIs.
- [ ] Make convenient wrappers for creating API handlers.
- [ ] Write the tests for all APIs.
- [x] Organize the code with packages
- [ ] Make docs with GoDoc
- [ ] Building a deployment process
