package userhandlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateHandler(t *testing.T) {
	var sentJson = []byte(`{
		"userId":10,
		"userName":"Anand",
		"userAge":20,
		"userLocation":"Bengaluru"
	
	}`)

	req, err := http.NewRequest("POST", "/v1/user/create", bytes.NewBuffer(sentJson))

	if err != nil {
		t.Fatal(err)
	}

	respRec := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateHandler)
	handler.ServeHTTP(respRec, req)

	if status := respRec.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	want := `{"statusCode":200,"error":"","message":"User Created","data":{"id":10,"userName":"Anand","userAge":20,"userLocation":"Bengaluru"}}`

	// respRec.Body.String() contains "\n" character, have to remove for comparison
	if strings.Compare(strings.Trim(respRec.Body.String(), "\n"), want) != 0 {
		t.Errorf("Handler returned wrong body : got %v want %v", strings.Trim(respRec.Body.String(), "\n"), want)
	}
}
