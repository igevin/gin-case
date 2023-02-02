package standard

import (
	"net/http"
	"testing"
)

type dummyHandler struct {
}

func (h dummyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("this is root"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Internal server error"))
	}
}

func TestServeHttp(t *testing.T) {
	_ = http.ListenAndServe(":8080", dummyHandler{})
}
