package httperrorutil_test

import (
	"errors"
	"testing"

	"github.com/onlineproducthouse/alchemypdf.api.httputils/helpers/asserterrmsg"
	"github.com/onlineproducthouse/alchemypdf.api.httputils/httperrorutil"
	"github.com/onlineproducthouse/alchemypdf.api.httputils/httpstatusutil"
)

func TestValidationErr(t *testing.T) {
	op := "TestValidationErr"
	msg := "Testing Validation Error"
	innerMsg := "Inner Testing Validation Error"
	originalErr := errors.New(innerMsg)

	err := httperrorutil.ValidationErr(msg, op, originalErr)

	if err.Kind() != httperrorutil.ValidationErrStr {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Error Kind", httperrorutil.ValidationErrStr, err.Kind()))
	}

	code, _ := httpstatusutil.BadRequest()
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

	err := httperrorutil.AuthErr(msg, op, originalErr)

	if err.Kind() != httperrorutil.AuthErrStr {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Error Kind", httperrorutil.AuthErrStr, err.Kind()))
	}
}

func TestForbiddenErr(t *testing.T) {
	op := "TestForbiddenErr"
	msg := "Testing ForbiddenErr Error"
	innerMsg := "Inner Testing ForbiddenErr Error"
	originalErr := errors.New(innerMsg)

	err := httperrorutil.ForbiddenErr(msg, op, originalErr)

	if err.Kind() != httperrorutil.ForbiddenErrStr {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Error Kind", httperrorutil.ForbiddenErrStr, err.Kind()))
	}
}

func TestNotFoundErr(t *testing.T) {
	op := "TestNotFoundErr"
	msg := "Testing NotFoundErr Error"
	innerMsg := "Inner Testing NotFoundErr Error"
	originalErr := errors.New(innerMsg)

	err := httperrorutil.NotFoundErr(msg, op, originalErr)

	if err.Kind() != httperrorutil.NotFoundErrStr {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Error Kind", httperrorutil.NotFoundErrStr, err.Kind()))
	}
}

func TestResorceLockedErr(t *testing.T) {
	op := "TestResorceLockedErr"
	msg := "Testing ResorceLockedErr Error"
	innerMsg := "Inner Testing ResorceLockedErr Error"
	originalErr := errors.New(innerMsg)

	err := httperrorutil.ResorceLockedErr(msg, op, originalErr)

	if err.Kind() != httperrorutil.ResorceLockedErrStr {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Error Kind", httperrorutil.ResorceLockedErrStr, err.Kind()))
	}
}

func TestUnknownErr(t *testing.T) {
	op := "TestUnknownErr"
	msg := "Testing UnknownErr Error"
	innerMsg := "Inner Testing UnknownErr Error"
	originalErr := errors.New(innerMsg)

	err := httperrorutil.UnknownErr(msg, op, originalErr)

	if err.Kind() != httperrorutil.UnknownErrStr {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Error Kind", httperrorutil.UnknownErrStr, err.Kind()))
	}
}

func TestNotImplementedErr(t *testing.T) {
	op := "TestNotImplementedErr"
	err := httperrorutil.NotImplementedErr(op)

	if err.Kind() != httperrorutil.NotImplementedErrStr {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Error Kind", httperrorutil.NotImplementedErrStr, err.Kind()))
	}
}

func TestDeprecated(t *testing.T) {
	op := "TestDeprecated"
	err := httperrorutil.Deprecated(op)

	if err.Kind() != httperrorutil.DeprecatedErrStr {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Error Kind", httperrorutil.DeprecatedErrStr, err.Kind()))
	}
}

func TestCatchErr(t *testing.T) {
	op := "TestCatchErr"
	msg := "Testing CatchErr Error"
	innerMsg := "Inner Testing CatchErr Error"
	originalErr := errors.New(innerMsg)

	err := httperrorutil.CatchErr(httperrorutil.UnknownErr(msg, op, originalErr), op)

	if err.Trace().InnerMessage != innerMsg {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Error Trace Message", msg, err.Trace().InnerMessage))
	}
}
