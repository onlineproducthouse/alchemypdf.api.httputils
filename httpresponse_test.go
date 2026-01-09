package alchemypdfapihttputils

import (
	"fmt"
	"testing"

	"alchemypdf.api.httputils/helpers/asserterrmsg"
)

func TestResponseDefault(t *testing.T) {
	_testCases := []Response{
		{"Ok", 200},
		{"Bad Request", 400},
		{"Unauthorized", 401},
		{"Forbidden", 403},
		{"Not Found", 404},
		{"Locked", 423},
		{"Internal Server Error", 500},
	}

	for _, expect := range _testCases {
		t.Run(fmt.Sprintf("HTTP Response %s", expect.Message), func(t *testing.T) {
			got := Default(expect.Message, expect.StatusCode)

			if got.Message != expect.Message {
				t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("Message", expect.Message, got.Message))
			}

			if got.StatusCode != expect.StatusCode {
				t.Errorf("%s", asserterrmsg.BuildAssertErrorMessage("StatusCode", expect.StatusCode, got.StatusCode))
			}
		})
	}
}

func BenchmarkResponseDefault(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Default("Ok", 200)
	}
}
