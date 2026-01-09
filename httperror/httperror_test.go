package httperror_test

import (
	"errors"
	"testing"

	"github.com/onlineproducthouse/alchemypdf.api.httputils/helpers/asserterrmsg"
	"github.com/onlineproducthouse/alchemypdf.api.httputils/httperror"
	"github.com/onlineproducthouse/alchemypdf.api.httputils/httpstatus"
)

func TestValidationErr(t *testing.T) {
	op := "TestValidationErr"
	msg := "Testing Validation Error"
	innerMsg := "Inner Testing Validation Error"
	originalErr := errors.New(innerMsg)

	err := httperror.ValidationErr(msg, op, originalErr)

	if err.Kind() != httperror.ValidationErrStr {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Error Kind", httperror.ValidationErrStr, err.Kind()))
	}

	code, _ := httpstatus.BadRequest()
	if err.StatusCode() != code {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Error StatusCode", code, err.StatusCode()))
	}

	if err.Error() != msg {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Error", msg, err.Error()))
	}

	if err.OriginalErr().Error() != originalErr.Error() {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("OriginalErr", originalErr, err.OriginalErr()))
	}

	if err.Trace().InnerMessage != innerMsg {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Trace InnerMessage", innerMsg, err.Trace().InnerMessage))
	}
}

func TestAuthErr(t *testing.T) {
	op := "TestAuthErr"
	msg := "Testing Auth Error"
	innerMsg := "Inner Testing Auth Error"
	originalErr := errors.New(innerMsg)

	err := httperror.AuthErr(msg, op, originalErr)

	if err.Kind() != httperror.AuthErrStr {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Error Kind", httperror.AuthErrStr, err.Kind()))
	}
}

func TestForbiddenErr(t *testing.T) {
	op := "TestForbiddenErr"
	msg := "Testing ForbiddenErr Error"
	innerMsg := "Inner Testing ForbiddenErr Error"
	originalErr := errors.New(innerMsg)

	err := httperror.ForbiddenErr(msg, op, originalErr)

	if err.Kind() != httperror.ForbiddenErrStr {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Error Kind", httperror.ForbiddenErrStr, err.Kind()))
	}
}

func TestNotFoundErr(t *testing.T) {
	op := "TestNotFoundErr"
	msg := "Testing NotFoundErr Error"
	innerMsg := "Inner Testing NotFoundErr Error"
	originalErr := errors.New(innerMsg)

	err := httperror.NotFoundErr(msg, op, originalErr)

	if err.Kind() != httperror.NotFoundErrStr {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Error Kind", httperror.NotFoundErrStr, err.Kind()))
	}
}

func TestResorceLockedErr(t *testing.T) {
	op := "TestResorceLockedErr"
	msg := "Testing ResorceLockedErr Error"
	innerMsg := "Inner Testing ResorceLockedErr Error"
	originalErr := errors.New(innerMsg)

	err := httperror.ResorceLockedErr(msg, op, originalErr)

	if err.Kind() != httperror.ResorceLockedErrStr {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Error Kind", httperror.ResorceLockedErrStr, err.Kind()))
	}
}

func TestUnknownErr(t *testing.T) {
	op := "TestUnknownErr"
	msg := "Testing UnknownErr Error"
	innerMsg := "Inner Testing UnknownErr Error"
	originalErr := errors.New(innerMsg)

	err := httperror.UnknownErr(msg, op, originalErr)

	if err.Kind() != httperror.UnknownErrStr {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Error Kind", httperror.UnknownErrStr, err.Kind()))
	}
}

func TestNotImplementedErr(t *testing.T) {
	op := "TestNotImplementedErr"
	err := httperror.NotImplementedErr(op)

	if err.Kind() != httperror.NotImplementedErrStr {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Error Kind", httperror.NotImplementedErrStr, err.Kind()))
	}
}

func TestDeprecated(t *testing.T) {
	op := "TestDeprecated"
	err := httperror.Deprecated(op)

	if err.Kind() != httperror.DeprecatedErrStr {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Error Kind", httperror.DeprecatedErrStr, err.Kind()))
	}
}

func TestCatchErr(t *testing.T) {
	op := "TestCatchErr"
	msg := "Testing CatchErr Error"
	innerMsg := "Inner Testing CatchErr Error"
	originalErr := errors.New(innerMsg)

	err := httperror.CatchErr(httperror.UnknownErr(msg, op, originalErr), op)

	if err.Trace().InnerMessage != innerMsg {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Error Trace Message", msg, err.Trace().InnerMessage))
	}
}
