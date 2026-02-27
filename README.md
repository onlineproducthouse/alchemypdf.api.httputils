# AlcheMyPDF API HTTP Utilities

## Installation

run:
```bash
go get github.com/onlineproducthouse/alchemypdf.api.httputils
```

## Usage

Import a module:

```go
import (
  "github.com/onlineproducthouse/alchemypdf.api.httputils/httperrorutil"
  "github.com/onlineproducthouse/alchemypdf.api.httputils/httprequtil"
  "github.com/onlineproducthouse/alchemypdf.api.httputils/httpresponseutil"
  "github.com/onlineproducthouse/alchemypdf.api.httputils/httpstatusutil"
)
```

Call a method:
```go
statusCode, statusCodeText := httpstatusutil.Ok()
```

## How it works

This package is used by the AlcheMyPDF API for handling http related operations. The package is split into four modules:
1. Error handling - [view example](./httperrorutil/httperrorutil_test.go)
2. HTTP response handling - [view example](./httpresponseutil/httpresponseutil_test.go)
3. HTTP status codes used - [view example](./httpstatusutil/httpstatusutil_test.go)

### Making HTTP requests (example)
```go
cfg := httprequtil.HTTPReqUtilCfg{
  Method: "GET",
  URL:    "https://example.org",
}

httpReq, httpReqErr := httprequtil.NewHTTPReqUtil(cfg)
if httpReqErr != nil {
  return nil, httperrorutil.CatchErr(httpReqErr, op)
}

httpReq.SetHeader("x-api-key", api.config.AlcheMyPDFAPIKey())

httpResp, httpRespBody, httpRespErr := httpReq.Execute()
if httpRespErr != nil {
  return nil, httperrorutil.CatchErr(httpRespErr, op)
}

if httpResp.StatusCode == 200 {
  var res []*Request
  if err := binary.Read(bytes.NewReader(httpRespBody), binary.BigEndian, &res); err != nil {
    return nil, httperrorutil.UnknownErr("failed to execute AlcheMyPDF API request", op, err)
  }

  return res, nil
} else {
  var res httpresponseutil.Response
  if err := binary.Read(bytes.NewReader(httpRespBody), binary.BigEndian, &res); err != nil {
    return nil, httperrorutil.UnknownErr("failed to execute AlcheMyPDF API request", op, err)
  }

  return nil, httperrorutil.UnknownErr(res.Message, op, nil)
}
```

## Unit tests

```bash
go test ./... -coverprofile=coverage-util.out
go tool cover -html=coverage-util.out
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)

