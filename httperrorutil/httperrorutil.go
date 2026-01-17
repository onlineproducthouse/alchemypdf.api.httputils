package httperrorutil

import "github.com/onlineproducthouse/alchemypdf.api.httputils/httpstatusutil"

const (
	NoErrStr             string = "Ok"                  // status 200
	ValidationErrStr     string = "ValidationError"     // status 400
	AuthErrStr           string = "AuthenticationError" // status 401
	ForbiddenErrStr      string = "ForbiddenError"      // status 403
	NotFoundErrStr       string = "NotFoundError"       // status 404
	ResorceLockedErrStr  string = "ResorceLockedError"  // status 423
	UnknownErrStr        string = "UnknownError"        // status 500
	NotImplementedErrStr string = "NotImplementedError" // status 501
	DeprecatedErrStr     string = "DeprecatedError"     // status 410
)

type IAppError interface {
	Kind() string
	StatusCode() int
	Error() string
	Trace() Trace
	OriginalErr() error
}

type AppError struct {
	name           string
	message        string
	op             string
	httpStatusCode int
	err            error
}

type Trace struct {
	InnerMessage string
	Ops          []string
}

func ValidationErr(msg, op string, err error) *AppError {
	code, _ := httpstatusutil.BadRequest()
	return &AppError{ValidationErrStr, msg, op, code, err}
}

func AuthErr(msg, op string, err error) *AppError {
	code, _ := httpstatusutil.Unauthorized()
	return &AppError{AuthErrStr, msg, op, code, err}
}

func ForbiddenErr(msg, op string, err error) *AppError {
	code, _ := httpstatusutil.Forbidden()
	return &AppError{ForbiddenErrStr, msg, op, code, err}
}

func NotFoundErr(msg, op string, err error) *AppError {
	code, text := httpstatusutil.NotFound()

	if msg != "" {
		text = msg
	}

	return &AppError{NotFoundErrStr, text, op, code, err}
}

func ResorceLockedErr(msg, op string, err error) *AppError {
	code, _ := httpstatusutil.Locked()
	return &AppError{ResorceLockedErrStr, msg, op, code, err}
}

func UnknownErr(msg, op string, err error) *AppError {
	code, _ := httpstatusutil.InternalServerError()
	return &AppError{UnknownErrStr, msg, op, code, err}
}

func NotImplementedErr(op string) *AppError {
	code, text := httpstatusutil.NotImplementedError()
	return &AppError{NotImplementedErrStr, text, op, code, nil}
}

func Deprecated(op string) *AppError {
	code, text := httpstatusutil.DeprecatedError()
	return &AppError{DeprecatedErrStr, text, op, code, nil}
}

func CatchErr(err IAppError, op string) *AppError {
	return &AppError{err.Kind(), err.Error(), op, err.StatusCode(), err}
}

func (a *AppError) Kind() string {
	return a.name
}

func (a *AppError) StatusCode() int {
	return a.httpStatusCode
}

func (a *AppError) Error() string {
	return a.message
}

func (a *AppError) OriginalErr() error {
	return a.err
}

func (a *AppError) Trace() Trace {
	return Trace{retrieveInnerMessage(a), retrieveOperations(a)}
}

func retrieveInnerMessage(a *AppError) string {
	if innerAppError, ok := a.err.(*AppError); ok {
		return retrieveInnerMessage(innerAppError)
	}

	if a.err == nil {
		return ""
	}

	return a.err.Error()
}

func retrieveOperations(a *AppError) []string {
	ops := []string{a.op}

	innerAppError, ok := a.err.(*AppError)
	if !ok {
		return ops
	}

	ops = append(ops, retrieveOperations(innerAppError)...)

	return ops
}
