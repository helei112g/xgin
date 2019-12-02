package benchmark

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkPing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/ping", nil)
		Engine.ServeHTTP(w, req)
	}
}
