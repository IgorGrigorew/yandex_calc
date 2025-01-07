package application

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	
)


func TestCalcHandler_ValidRequest(t *testing.T) {
	

	requestBody := Request{Expression: "2+2"}
	body, _ := json.Marshal(requestBody)

	req, err := http.NewRequest("POST", "/api/v1/calculate", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CalcHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var okResponse OkRequest
	if err := json.Unmarshal(rr.Body.Bytes(), &okResponse); err != nil {
		t.Fatalf("could not unmarshal response: %v", err)
	}

	if okResponse.Result != 4 {
		t.Errorf("handler returned unexpected body: got %v want %v", okResponse.Result, 4)
	}
}

// Test for invalid calculation
func TestCalcHandler_InvalidRequest(t *testing.T) {


	requestBody := Request{Expression: "invalid"}
	body, _ := json.Marshal(requestBody)

	req, err := http.NewRequest("POST", "/api/v1/calculate", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CalcHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnprocessableEntity {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnprocessableEntity)
	}

	var badResponse BadRequest
	if err := json.Unmarshal(rr.Body.Bytes(), &badResponse); err != nil {
		t.Fatalf("could not unmarshal response: %v", err)
	}

	expectedErrorMsg := "Expression is not valid"
	if badResponse.Error != expectedErrorMsg {
		t.Errorf("handler returned unexpected body: got %v want %v", badResponse.Error, expectedErrorMsg)
	}
}
