package middleware_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/EricDriussi/api-pet-hotel-go/internal/infrastructure/server/middleware"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMiddleware_PrintsDiscInferno(t *testing.T) {
	gin.SetMode(gin.TestMode)
	engine := gin.New()
	engine.Use(middleware.DiscoInferno)

	outputRecorder := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	httpRecorder := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/irrelevant", nil)
	require.NoError(t, err)

	engine.ServeHTTP(httpRecorder, req)

	require.NoError(t, w.Close())
	got, _ := ioutil.ReadAll(r)
	os.Stdout = outputRecorder

	assert.Contains(t, string(got), " 🔥 ")
	assert.Contains(t, string(got), "Burn baby burn")
	assert.Contains(t, string(got), "Disco inferno")
}
