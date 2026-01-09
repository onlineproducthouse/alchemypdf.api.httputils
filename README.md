# AlcheMyPDF API HTTP Utilities

## Installation

run:
```bash
go get github.com/onlineproducthouse/alchemypdf.api.httputils
```

## Usage

Import the package:

```go
import alchemypdfapihttputils "github.com/onlineproducthouse/alchemypdf.api.httputils"
```

Call a method:
```go
statusCode, statusCodeText := alchemypdfapihttputils.Ok()
```

## Unit tests

```bash
go test ./... -coverprofile=coverage-util.out
go tool cover -html=coverage-util.out
```

## About

This package is used by the AlcheMyPDF API for handling http related operations. The package is split into four areas:
1. Error handling
2. Making HTTP requests
3. HTTP response handling
4. HTTP status codes used
