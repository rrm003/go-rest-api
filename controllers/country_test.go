package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/countries", FetchCountries)
	return router
}

func TestFetchCountries(t *testing.T) {
	defer gock.Off()

	router := setupRouter()

	t.Run("Successful Fetch", func(t *testing.T) {
		gock.New("https://api.first.org").
			Get("/data/v1/countries").
			Reply(200).
			JSON(map[string]interface{}{
				"data": map[string]interface{}{
					"USA": map[string]string{"country": "Burundi"},
					"CAN": map[string]string{"country": "Afghanistan"},
				},
			})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/countries", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "Burundi")
		assert.Contains(t, w.Body.String(), "Afghanistan")
	})
}
