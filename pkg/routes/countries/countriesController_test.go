package countries

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// func TestAllCountriesRoute(t *testing.T) {
// 	req := httptest.NewRequest("GET", "localhost:8080/", nil)
// 	w := httptest.NewRecorder()
// 	All(w, req)
// }

func TestSingleCountryRoute(t *testing.T) {
	req := httptest.NewRequest("GET", "localhost:8080/", nil)
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(FindOne)
	handler.ServeHTTP(w, req)
	fmt.Println("This is too much")
}
