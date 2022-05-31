# hash-app-go
Small App to Hash Password Technical Exercise

## Assumptions

The following assumptions are being made:

* Since the exercise clearly states that only standard libraries should be used, I am not using a database connection since that will require the use of external driver libraries. Otherwise, I would have used a DB to manage the password hashes
* I'm also assuming that not web framework can be used.
* I'm also assuming that all unit tests for this service will run through a CI/CD pipeline before being built and deployed. 

## Requirements

## Go Tools
Follow the [Download and Install](https://go.dev/doc/install) instructions to get Go installed in your machine.

## Setting and Running Project

Checkout project from Github

`git clone https://github.com/davicho01/hash-app-go`

Go to `app` directory inside project `cd hash-app-go/app` directory

Run `go run main.go`

## Docker build

This project already comes integrate it with dockerfile. To build docker please install [docker](https://docs.docker.com/get-docker/) on your machine

### To build and run docker container:

Go to director `hash-app-go`

Build container `docker build -t hash-app-go .`

Run container `docker run --rm -it -p 8080:8080 hash-app-go`

## Unit Testing

To run all Unit test:

Go to `hash-app-go/app` directory

Run `go test ./tests/... -v`

## Testing API

Following is a list of endpoints that can be used to test this application.

#### Post password
```curl
curl --location --request POST 'http://localhost:8080/hash' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'password=angryMonkey'
```

#### Get Hash
```curl
curl --location --request GET 'http://localhost:8080/hash/1'
```

#### Get Stats
```curl
curl --location --request GET 'http://localhost:8080/stats'
```

#### Shutdown
````curl
curl --location --request GET 'http://localhost:8080/shutdown'
````


