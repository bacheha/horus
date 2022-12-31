package middlewares

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestValidateObjectID(t *testing.T) {
	// target
	req, _ := http.NewRequest(http.MethodGet, "/12345", nil)
	rr := httptest.NewRecorder()
	mux := chi.NewRouter()
	mux.Use(ValidateObjectID("id"))
	mux.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello %s", chi.URLParam(r, "id"))
	})
	mux.ServeHTTP(rr, req)

	// assert
	res := rr.Result()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("status code is not %d, got %d", http.StatusBadRequest, res.StatusCode)
	}
}
