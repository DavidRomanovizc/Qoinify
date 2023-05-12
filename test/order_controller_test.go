package test

import (
	"github.com/DavidRomanovizc/Qoinify/internal/api/controllers"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPlaceOrderIntegration(t *testing.T) {
	router := gin.Default()
	router.GET("/place-order", controllers.GetPlaceOrder)
	server := httptest.NewServer(router)
	defer server.Close()

	resp, err := http.Get(server.URL + "/place-order")
	if err != nil {
		t.Errorf("Error making HTTP request: %s", err.Error())
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, but got %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error reading response body: %s", err.Error())
	}
	expected := `{"message":"pong"}`
	if string(body) != expected {
		t.Errorf("Expected response body '%s', but got '%s'", expected, string(body))
	}
}
