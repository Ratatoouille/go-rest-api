package api

import (
	"net/http"
	"net/http/httptest"
	"restApi/internal/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApiServer_HandleHello(t *testing.T) {
	s := NewAPI(config.NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	s.handleHello().ServeHTTP(rec, req)

	assert.Equal(t, rec.Body.String(), "hello")
}
