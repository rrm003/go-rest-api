// controllers/user_controller_test.go
package controllers

import (
	"errors"
	"go-rest-api/models"
	svcMock "go-rest-api/services/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func setupTest() (*gin.Engine, *svcMock.MockUserService, *gomock.Controller) {
	ctrl := gomock.NewController(nil)
	mockUserService := svcMock.NewMockUserService(ctrl)
	userController := NewUserController(mockUserService)

	router := gin.Default()
	router.POST("/signup", userController.SignUp)
	router.POST("/login", userController.Login)
	router.GET("/users", userController.GetUsers)
	router.GET("/users/:id", userController.GetUser)
	router.PUT("/users/:id", userController.UpdateUser)
	router.DELETE("/users/:id", userController.DeleteUser)

	return router, mockUserService, ctrl
}

func TestSignUp(t *testing.T) {
	router, mockUserService, ctrl := setupTest()
	defer ctrl.Finish()

	user := models.User{
		Username: "testuser",
		Password: "password",
		Country:  "usa",
	}

	mockUserService.EXPECT().SignUp(user).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/signup", strings.NewReader(`{"username":"testuser","password":"password", "country":"usa"}`))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "testuser")
}

func TestSignUp_Fail_NoPassword(t *testing.T) {
	router, mockUserService, ctrl := setupTest()
	defer ctrl.Finish()

	user := models.User{
		Username: "testuser",
		Country:  "usa",
	}

	mockUserService.EXPECT().SignUp(user).Return(nil).Times(0)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/signup", strings.NewReader(`{"username":"testuser","country":"usa"}`))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "required fields [country, password]")
}

func TestSignUp_Fail_NoCountry(t *testing.T) {
	router, mockUserService, ctrl := setupTest()
	defer ctrl.Finish()

	user := models.User{
		Username: "testuser",
		Password: "asopa#010",
	}

	mockUserService.EXPECT().SignUp(user).Return(nil).Times(0)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/signup", strings.NewReader(`{"username":"testuser","password":"asopa#010"}`))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "required fields [country, password]")
}

func TestSignUp_Fail_DBErr(t *testing.T) {
	router, mockUserService, ctrl := setupTest()
	defer ctrl.Finish()

	user := models.User{
		Username: "testuser",
		Password: "password",
		Country:  "usa",
	}

	mockUserService.EXPECT().SignUp(user).Return(errors.New("failed to fetch record"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/signup", strings.NewReader(`{"username":"testuser","password":"password", "country":"usa"}`))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "failed to fetch record")
}

func TestLogin(t *testing.T) {
	router, mockUserService, ctrl := setupTest()
	defer ctrl.Finish()

	username := "testuser"
	password := "password"
	token := "mocked-jwt-token"

	mockUserService.EXPECT().Login(username, password).Return(token, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(`{"username":"testuser","password":"password"}`))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), token)
}

func TestGetUsers(t *testing.T) {
	router, mockUserService, ctrl := setupTest()
	defer ctrl.Finish()

	users := []models.User{
		{Username: "user1", Password: "password1"},
		{Username: "user2", Password: "password2"},
	}

	mockUserService.EXPECT().GetUsers().Return(users, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "user1")
	assert.Contains(t, w.Body.String(), "user2")
}

func TestGetUser(t *testing.T) {
	router, mockUserService, ctrl := setupTest()
	defer ctrl.Finish()

	user := models.User{Username: "testuser", Password: "password"}

	mockUserService.EXPECT().GetUser("1").Return(user, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "testuser")
}

func TestUpdateUser(t *testing.T) {
	router, mockUserService, ctrl := setupTest()
	defer ctrl.Finish()

	user := models.User{Username: "updateduser", Password: "newpassword"}

	mockUserService.EXPECT().UpdateUser("1", user).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/users/1", strings.NewReader(`{"username":"updateduser","password":"newpassword"}`))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "updateduser")
}

func TestDeleteUser(t *testing.T) {
	router, mockUserService, ctrl := setupTest()
	defer ctrl.Finish()

	mockUserService.EXPECT().DeleteUser("1").Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/users/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "User deleted")
}
