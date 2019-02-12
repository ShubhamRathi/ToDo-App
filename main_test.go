package main

import (
	"net/http"
	"net/http/httptest"
    "testing"
    "github.com/gorilla/mux"
    "github.com/stretchr/testify/assert"
    // "github.com/DATA-DOG/go-sqlmock"
)

func Router() *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/items", AllItems).Methods("GET")
    router.HandleFunc("/items/{text}", AddItem).Methods("POST")
    router.HandleFunc("/items/{text}/{update}", UpdateItem).Methods("PUT")
    router.HandleFunc("/items/{text}", UpdateItem).Methods("DELETE")
    return router
}

func TestHandleRequestsSuccess(t *testing.T) {
    request, _ := http.NewRequest("GET", "/items", nil)
    response := httptest.NewRecorder()
    Router().ServeHTTP(response, request)
    assert.Equal(t, 200, response.Code, "OK response is expected")

    request1, _ := http.NewRequest("POST", "/items/Hi", nil)
    response1 := httptest.NewRecorder()
    Router().ServeHTTP(response1, request1)
    assert.Equal(t, 200, response1.Code, "OK response is expected")

    request2, _ := http.NewRequest("PUT", "/items/Hi/Bye", nil)
    response2 := httptest.NewRecorder()
    Router().ServeHTTP(response2, request2)
    assert.Equal(t, 200, response2.Code, "OK response is expected")

    request3, _ := http.NewRequest("DELETE", "/items/Bye", nil)
    response3 := httptest.NewRecorder()
    Router().ServeHTTP(response3, request3)
    assert.Equal(t, 200, response3.Code, "OK response is expected")

}

// func TestHandleRequestsFailure(t *testing.T) { // Needs help

//     request2, _ := http.NewRequest("PUT", "/items/SampleNotPresent/Present", nil)
//     response2 := httptest.NewRecorder()
//     Router().ServeHTTP(response2, request2)
//     assert.NotEqual(t, 200, response2.Code, "OK response is not expected")

//     request3, _ := http.NewRequest("DELETE", "/items/Present", nil)
//     response3 := httptest.NewRecorder()
//     Router().ServeHTTP(response3, request3)
//     assert.NotEqual(t, 200, response3.Code, "OK response is not expected")

// }


