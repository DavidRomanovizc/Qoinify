package test

import (
	"net/http"
	"testing"
)

func TestGetPlaceOrderEndpoint(t *testing.T) {

	responses := make(chan *http.Response)

	for i := 0; i < 10; i++ {
		go func() {
			resp, err := http.Get("http://0.0.0.0:8080/api/ping")
			if err != nil {
				t.Errorf("error sending GET request: %v", err)
			}
			responses <- resp
		}()
	}

	for i := 0; i < 10; i++ {
		resp := <-responses
		if resp.StatusCode != 200 {
			t.Errorf("unexpected status code: %v", resp.StatusCode)
		}
	}
}
