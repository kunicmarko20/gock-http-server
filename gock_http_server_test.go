package main

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers(t *testing.T) {
	router := newRouter()
	addMock(t, router)

	recorder := getMock(t, router)

	assertStatusCodeEquals(t, http.StatusOK, recorder.Code)

	data := getDataFromRecorder(recorder, t)

	if data["example"] != "response" {
		t.Errorf("Expected to find \"{\"example\": \"response\"}\" instead found: %s", data)
	}

	reset(t, router)

	recorder = getMock(t, router)

	assertStatusCodeEquals(t, http.StatusBadRequest, recorder.Code)
}

func addMock(t *testing.T, router *mux.Router) {
	request, err := http.NewRequest(
		"POST",
		"/api/mock/test",
		bytes.NewBuffer(createAddMockBody()),
	)

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	assertStatusCodeEquals(t, http.StatusNoContent, recorder.Code)
}

func getMock(t *testing.T, router *mux.Router) *httptest.ResponseRecorder {
	request, err := http.NewRequest(
		"GET",
		"/mock/foo/bar",
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	return recorder
}

func assertStatusCodeEquals(t *testing.T, expected int, actual int) {
	if expected != actual {
		t.Errorf(
			"status code was expeted to be %d, instead received %d",
			expected,
			actual,
		)
	}
}

func reset(t *testing.T, router *mux.Router) {
	request, err := http.NewRequest(
		"POST",
		"/api/reset",
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	assertStatusCodeEquals(t, http.StatusNoContent, recorder.Code)
}

func getDataFromRecorder(recorder *httptest.ResponseRecorder, t *testing.T) map[string]string {
	decoder := json.NewDecoder(recorder.Body)

	var data map[string]string

	err := decoder.Decode(&data)

	if err != nil {
		t.Fatal(err)
	}

	return data
}

func createAddMockBody() []byte {
	return []byte(`
{
  "matchRule": {
    "type": "allOf",
    "rules": [
      {
        "type": "pathEquals", 
        "value": "foo/bar" 
      },
      {
        "type": "anyOf",
        "rules": [
          {
            "type": "methodEquals",
            "value": "POST"
          },
          {
            "type": "methodEquals",
            "value": "GET"
          }
        ]       
      }
    ]
  },
  "response": {
    "statusCode": 200,
    "headers": {
      "Content-Type": "application/json; charset=utf-8"
    },
    "content": "{\"example\": \"response\"}"
  }  
}
`)
}
