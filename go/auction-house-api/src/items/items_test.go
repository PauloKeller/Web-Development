package items

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test(t *testing.T) {
	items := "[{'id': '3e9192ef-a1aa-4744-8acc-a6696a3f18ce','name': 'alo','description': 'testamos'},{'id': 'f6886870-ece2-4e7d-a539-2e8ba63faf7f','name': 'Item 1','description': 'a big text here!'}]"
	item := []byte("{ 'name':  'Item 1','description':'a big text here!'}")

	tests := []struct {
		name           string
		in             *http.Request
		out            *httptest.ResponseRecorder
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "retrieve all items",
			in:             httptest.NewRequest("GET", "/items", nil),
			out:            httptest.NewRecorder(),
			expectedStatus: http.StatusOK,
			expectedBody:   items,
		},
		{
			name:           "create item",
			in:             httptest.NewRequest("POST", "/items", bytes.NewBuffer(item)),
			out:            httptest.NewRecorder(),
			expectedStatus: http.StatusOK,
			expectedBody:   items,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			if test.name == "retrieve all items" {
				FetchItems(test.out, test.in)
			}

			if test.name == "create item" {
				CreateItem(test.out, test.in)
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
