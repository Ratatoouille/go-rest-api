package api

import (
	"net/http"
	"net/http/httptest"
	"restApi/internal/store/teststore"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer_UsersCreate(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/users", nil)

	s := newServer(teststore.New())

	s.ServeHTTP(rec, req)

	assert.Equal(t, rec.Code, http.StatusOK)
}
