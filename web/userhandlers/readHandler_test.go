package userhandlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestReadHandler(t *testing.T) {

	t.Run("GET User found", func(t *testing.T) {

		req, err := http.NewRequest("GET", "/v1/user/read/", nil)

		if err != nil {
			t.Fatal(err)
		}
		q := req.URL.Query()
		q.Add("userId", "10")
		req.URL.RawQuery = q.Encode()

		respRec := httptest.NewRecorder()
		handler := http.HandlerFunc(ReadHandler)

		handler.ServeHTTP(respRec, req)

		if status := respRec.Code; status != http.StatusOK {
			t.Errorf("Handler returned wrong status code : got %v want %v", status, http.StatusOK)
		}
		want := `{"statusCode":200,"error":"","message":"User Found","data":{"id":10,"userName":"Anand","userAge":20,"userLocation":"Bengaluru"}}`

		if strings.Compare(strings.Trim(respRec.Body.String(), "\n"), want) != 0 {
			t.Errorf("Handler returned wrong body : got %v want %v", strings.Trim(respRec.Body.String(), "\n"), want)
		}
	})
	t.Run("GET User Not found", func(t *testing.T) {

		req, err := http.NewRequest("GET", "/v1/user/read/", nil)

		if err != nil {
			t.Fatal(err)
		}
		q := req.URL.Query()
		q.Add("userId", "20")
		req.URL.RawQuery = q.Encode()

		respRec := httptest.NewRecorder()
		handler := http.HandlerFunc(ReadHandler)

		handler.ServeHTTP(respRec, req)

		if status := respRec.Code; status != http.StatusOK {
			t.Errorf("Handler returned wrong status code : got %v want %v", status, http.StatusOK)
		}

		want := `{"statusCode":200,"error":"","message":"User Not Found","data":null}`

		if strings.Compare(strings.Trim(respRec.Body.String(), "\n"), want) != 0 {
			t.Errorf("Handler returned wrong body : got %v want %v", strings.Trim(respRec.Body.String(), "\n"), want)
		}
	})

}
