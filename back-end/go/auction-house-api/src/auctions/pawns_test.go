package auctions

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test(t *testing.T) {
	auctions := ""
	auction := []byte("")

	tests := []struct {
		name           string
		in             *http.Request
		out            *httptest.ResponseRecorder
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "retrieve all auctions",
			in:             httptest.NewRequest("GET", "/auctions", nil),
			out:            httptest.NewRecorder(),
			expectedStatus: http.StatusOK,
			expectedBody:   pawns,
		},
		{
			name:           "create auction",
			in:             httptest.NewRequest("POST", "/auctions", bytes.NewBuffer(auction)),
			out:            httptest.NewRecorder(),
			expectedStatus: http.StatusOK,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			if test.name == "retrieve all auctions" {
				FetchPawns(test.out, test.in)
			}

			if test.name == "create auction" {
				CreatePawn(test.out, test.in)
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
