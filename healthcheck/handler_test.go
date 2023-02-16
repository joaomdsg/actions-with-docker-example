package healthcheck_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joaomdsg/actions-with-docker-example/healthcheck"
	"github.com/stretchr/testify/assert"
)

func TestHealthcheckHandlerFn(t *testing.T) {
	// given
	e := gin.Default()
	handlerFn := healthcheck.NewHandlerFn()
	e.GET("/healthcheck", handlerFn)
	req, _ := http.NewRequest("GET", "/healthcheck", nil)
	w := httptest.NewRecorder()

	// when
	e.ServeHTTP(w, req)

	// then
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"status":"healthy"}`, w.Body.String())

}
