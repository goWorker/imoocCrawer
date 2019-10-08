package test

import (
	"GinHello/initRouter"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexGetRouter(t *testing.T){
	router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	req,_ := http.NewRequest(http.MethodGet,"/",nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK,w.Code)
	assert.Equal(t,"hello gin get method",w.Body.String())
}

func TestIndexPostRouter(t *testing.T) {
	router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "hello gin post method", w.Body.String())
}

var router *gin.Engine

func init() {
	gin.SetMode(gin.TestMode)
	router = initRouter.SetupRouter()
}

func TestIndexHtml(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "hello gin get method", "返回的HTML页面中应该包含 hello gin get method")
}
