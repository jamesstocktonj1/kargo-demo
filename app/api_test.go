package app

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGreet(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router, err := NewRouter()
	assert.NoError(t, err)

	tests := []struct {
		name       string
		body       string
		wantStatus int
		wantBody   string
	}{
		{"valid name", `{"name":"James"}`, http.StatusOK, `{"message":"Hello, James!"}`},
		{"missing name", `{}`, http.StatusBadRequest, `{"error":"name is required"}`},
		{"blank name", `{"name":"   "}`, http.StatusBadRequest, `{"error":"name is required"}`},
		{"invalid json", `not json`, http.StatusBadRequest, `{"error":"invalid JSON body"}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/greet", strings.NewReader(tt.body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
			assert.Equal(t, tt.wantBody, strings.TrimSpace(w.Body.String()))
		})
	}
}
