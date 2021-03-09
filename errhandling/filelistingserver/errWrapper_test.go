package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}

type testuserError string

func (e testuserError) Error() string {
	return e.Message()
}

func (e testuserError) Message() string {
	return string(e)
}

func errUserError(writer http.ResponseWriter, request *http.Request) error {
	return testuserError("user error")
}

func errNotFound(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}

func errNotPremited(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}

func errUnknown(writer http.ResponseWriter, request *http.Request) error {
	return errors.New("unknown")
}

func noError(writer http.ResponseWriter, request *http.Request) error {
	fmt.Fprintln(writer, "no error")
	return nil
}

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNotFound, 404, "Not Found"},
	{errNotPremited, 403, "Forbidden"},
	{errUnknown, 500, "Internal Server Error"},
	{noError, 200, "no error"},
}

func TestErrWrapper(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "http://www.baidu.com", nil)
		f(response, request)

		verifyResponse(response.Result(), tt.code, tt.message, t)
	}
}

func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		response, _ := http.Get(server.URL)

		verifyResponse(response, tt.code, tt.message, t)
	}
}

func verifyResponse(response *http.Response, expectedCode int, expectedMessage string, t *testing.T) {
	b, _ := ioutil.ReadAll(response.Body)
	body := strings.Trim(string(b), "\n")
	if response.StatusCode != expectedCode || body != expectedMessage {
		t.Errorf("expect(%d %s); get(%d, %s)", expectedCode, expectedMessage, response.StatusCode, body)
	}
}
