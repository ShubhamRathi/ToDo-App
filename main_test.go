package main

import (
	"net/http"
	"net/http/httptest"
    "testing"
    "github.com/gorilla/mux"
    "github.com/stretchr/testify/assert"
    "github.com/DATA-DOG/go-sqlmock"
)

func Router() *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/items", AllItems).Methods("GET")
    return router
}

func TestHandleRequests(t *testing.T) {
	// Whats the DRY version of this?
    request, _ := http.NewRequest("GET", "/", nil)
    response := httptest.NewRecorder()
    Router().ServeHTTP(response, request)
    assert.Equal(t, 200, response.Code, "OK response is expected")

    request1, _ := http.NewRequest("POST", "/India", nil)
    response1 := httptest.NewRecorder()
    Router().ServeHTTP(response1, request1)
    assert.Equal(t, 200, response1.Code, "OK response is expected")
}
