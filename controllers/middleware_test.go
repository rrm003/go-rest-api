package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {
	router := gin.New()
	router.Use(AuthMiddleware())
	router.GET("/test", func(c *gin.Context) {
		username := c.MustGet("username").(string)
		c.JSON(http.StatusOK, gin.H{"username": username})
	})

	t.Run("No Authorization Header", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Contains(t, w.Body.String(), "Request does not contain an access token")
	})

	t.Run("Invalid Token - Incorrect Signing Method", func(t *testing.T) {
		token := jwt.NewWithClaims(jwt.SigningMethodHS512, &Claims{
			Username: "testuser",
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour).Unix(),
			},
		})
		tokenStr, _ := token.SignedString([]byte("my_secret_key"))

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", tokenStr)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Contains(t, w.Body.String(), "Invalid token")
	})

	t.Run("Invalid Token - Tampered Token", func(t *testing.T) {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
			Username: "testuser",
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour).Unix(),
			},
		})
		tokenStr, _ := token.SignedString([]byte("wrong_secret_key"))

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", tokenStr)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Contains(t, w.Body.String(), "Invalid token")
	})

	t.Run("Valid Token", func(t *testing.T) {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
			Username: "testuser",
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour).Unix(),
			},
		})
		tokenStr, _ := token.SignedString(jwtKey)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", tokenStr)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "testuser")
	})
}
