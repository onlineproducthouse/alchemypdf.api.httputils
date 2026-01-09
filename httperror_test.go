package alchemypdfapihttputils

import (
	"errors"
	"testing"

	"alchemypdf.api.httputils/helpers/asserterrmsg"
)

func TestValidationErr(t *testing.T) {
	op := "TestValidationErr"
	msg := "Testing Validation Error"
	innerMsg := "Inner Testing Validation Error"
	originalErr := errors.New(innerMsg)

	err := ValidationErr(msg, op, originalErr)

	if err.Kind() != ValidationErrStr {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Error Kind", ValidationErrStr, err.Kind()))
	}

	code, _ := BadRequest()
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

	err := AuthErr(msg, op, originalErr)

	if err.Kind() != AuthErrStr {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Error Kind", AuthErrStr, err.Kind()))
	}
}

func TestForbiddenErr(t *testing.T) {
	op := "TestForbiddenErr"
	msg := "Testing ForbiddenErr Error"
	innerMsg := "Inner Testing ForbiddenErr Error"
	originalErr := errors.New(innerMsg)

	err := ForbiddenErr(msg, op, originalErr)

	if err.Kind() != ForbiddenErrStr {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Error Kind", ForbiddenErrStr, err.Kind()))
	}
}

func TestNotFoundErr(t *testing.T) {
	op := "TestNotFoundErr"
	msg := "Testing NotFoundErr Error"
	innerMsg := "Inner Testing NotFoundErr Error"
	originalErr := errors.New(innerMsg)

	err := NotFoundErr(msg, op, originalErr)

	if err.Kind() != NotFoundErrStr {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Error Kind", NotFoundErrStr, err.Kind()))
	}
}

func TestResorceLockedErr(t *testing.T) {
	op := "TestResorceLockedErr"
	msg := "Testing ResorceLockedErr Error"
	innerMsg := "Inner Testing ResorceLockedErr Error"
	originalErr := errors.New(innerMsg)

	err := ResorceLockedErr(msg, op, originalErr)

	if err.Kind() != ResorceLockedErrStr {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Error Kind", ResorceLockedErrStr, err.Kind()))
	}
}

func TestUnknownErr(t *testing.T) {
	op := "TestUnknownErr"
	msg := "Testing UnknownErr Error"
	innerMsg := "Inner Testing UnknownErr Error"
	originalErr := errors.New(innerMsg)

	err := UnknownErr(msg, op, originalErr)

	if err.Kind() != UnknownErrStr {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Error Kind", UnknownErrStr, err.Kind()))
	}
}

func TestNotImplementedErr(t *testing.T) {
	op := "TestNotImplementedErr"
	err := NotImplementedErr(op)

	if err.Kind() != NotImplementedErrStr {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Error Kind", NotImplementedErrStr, err.Kind()))
	}
}

func TestDeprecated(t *testing.T) {
	op := "TestDeprecated"
	err := Deprecated(op)

	if err.Kind() != DeprecatedErrStr {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Error Kind", DeprecatedErrStr, err.Kind()))
	}
}

func TestCatchErr(t *testing.T) {
	op := "TestCatchErr"
	msg := "Testing CatchErr Error"
	innerMsg := "Inner Testing CatchErr Error"
	originalErr := errors.New(innerMsg)

	err := CatchErr(UnknownErr(msg, op, originalErr), op)

	if err.Trace().InnerMessage != innerMsg {
		t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Error Trace Message", msg, err.Trace().InnerMessage))
	}
}
