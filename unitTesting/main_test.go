package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIntegration(t *testing.T) {
	go main()
	// time.Sleep(time.Second * 5)

	//newStock := Stock{
	//	Symbol: "ZopSmart",
	//	Price:  1500.00,
	//	Volume: 100,
	//}

	var stockTestCases = []struct {
		desc       string
		method     string
		endpoint   string
		statusCode int
		body       []byte
	}{
		{"get success", http.MethodGet, "stock/12345", http.StatusOK, nil},
		{"create success", http.MethodPost, "stock", http.StatusCreated, []byte(`{"symbol":"ZopSmart","price":1500.00,"volume":100}`)},
		{"get unknown endpoint", http.MethodGet, "unknown", http.StatusNotFound, nil},
		{"get invalid endpoint", http.MethodGet, "stock/id", http.StatusNotFound, nil},
		{"unregistered route", http.MethodPut, "stock", http.StatusMethodNotAllowed, nil},
		{"delete success", http.MethodDelete, "stock/67890", http.StatusNoContent, nil},
	}

	jsonStock, err := json.Marshal(stockTestCases)
	if err != nil {
		t.Errorf("Error marshaling stock to JSON: %v", err)
	}
	req, err := http.NewRequest("POST", "/add", strings.NewReader(string(jsonStock)))
	if err != nil {
		t.Errorf("Error creating HTTP request: %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.ServeHTTP(w, r)
	})
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}
