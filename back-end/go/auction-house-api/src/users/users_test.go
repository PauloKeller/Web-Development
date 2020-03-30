package users

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test(t *testing.T) {
	users := ""
	user := []byte("")

	tests := []struct {
		name           string
		in             *http.Request
		out            *httptest.ResponseRecorder
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "retrieve all users",
			in:             httptest.NewRequest("GET", "/users", nil),
			out:            httptest.NewRecorder(),
			expectedStatus: http.StatusOK,
			expectedBody:   users,
		},
		{
			name:           "create item",
			in:             httptest.NewRequest("POST", "/users", bytes.NewBuffer(user)),
			out:            httptest.NewRecorder(),
			expectedStatus: http.StatusOK,
			expectedBody:   user,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			if test.name == "retrieve all users" {
				FetchUsers(test.out, test.in)
			}

			if test.name == "create user" {
				CreateUser(test.out, test.in)
			}

			if test.out.Code != test.expectedStatus {
				t.Logf("expected: %d\ngot: %d\n", test.expectedStatus, test.out.Code)
				t.Fail()
			}

			body := test.out.Body.String()
			if body != test.expectedBody {
				t.Logf("expected: %s\ngot: %s\n", test.expectedBody, body)
				t.Fail()
			}
		})
	}
}
