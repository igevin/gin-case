package standard

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestServeHttpNoResponse(t *testing.T) {
	// 这样就可以启动监听，但不会有任何请求响应
	err := http.ListenAndServe(":8080", nil)
	assert.NoError(t, err)
}

func TestServeHttpDummyResponse(t *testing.T) {
	// "/"会匹配所有URL path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("default response"))
	})

	err := http.ListenAndServe(":8080", nil)
	assert.NoError(t, err)
}

func TestServeHttpResponseV1(t *testing.T) {
	// "/"会匹配所有URL path，所以我们可以通过 http.Request 对象，分别匹配不同的URL，并做不同的处理
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/version" && r.Method == http.MethodPost {
			_, _ = w.Write([]byte("v1.0.0"))
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("default response"))
	})

	err := http.ListenAndServe(":8080", nil)
	assert.NoError(t, err)
}

func TestServeHttpResponseV2(t *testing.T) {
	// 这个是兜底的
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("default response"))
	})

	// 这个精确匹配 /version
	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("1.0.0"))
	})

	// 这个匹配 /api 前缀的
	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("api page"))
	})

	// 这个匹配 /admin 前缀的页面
	http.HandleFunc("/admin/", func(w http.ResponseWriter, r *http.Request) {
		s, err := json.Marshal(map[string]any{"success": true, "msg": "Welcome to admin page"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Internal Server error"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Add("key", "value")
		// 先设置 header，再write header
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(s)
	})

	err := http.ListenAndServe(":8080", nil)
	assert.NoError(t, err)
}
