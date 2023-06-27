package helper

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/httptest"
)

func CreateTestServer(request *http.Request, route func(router *gin.Engine)) (int, []byte) {
	gin.SetMode(gin.TestMode)
	responseRecorder := httptest.NewRecorder()
	_, router := gin.CreateTestContext(responseRecorder)

	route(router)
	router.ServeHTTP(responseRecorder, request)

	response := responseRecorder.Result()
	body, _ := io.ReadAll(response.Body)
	defer func() {
		_ = response.Body.Close()
	}()

	return response.StatusCode, body
}
