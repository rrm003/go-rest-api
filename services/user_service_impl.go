package services

import (
	"errors"
	"fmt"
	"go-rest-api/database"
	"go-rest-api/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

var jwtKey = []byte("secret_key")

type userService struct {
	db database.Database
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func NewUserService(db database.Database) UserService {
	return &userService{db: db}
}

func (s *userService) SignUp(user models.User) error {
	return s.db.Create(&user).Error
}

func (s *userService) Login(username, password string) (string, error) {
	var user models.User
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", errors.New("invalid credentials: username")
		}
		return "", err
	}

	// Validate password (simplified for this example, use proper hashing in production)
	if user.Password != password {
		return "", errors.New("invalid credentials: password")
	}

	// Create JWT token
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *userService) GetUsers() ([]models.User, error) {
	var users []models.User
	if err := s.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (s *userService) GetUser(id string) (models.User, error) {
	var user models.User
	if err := s.db.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (s *userService) UpdateUser(id string, user models.User) error {
	existing := models.User{}
	if err := s.db.First(&existing, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("user with ID %s not found", id)
		}
		return err
	}

	existing.Country = user.Country
	existing.Password = user.Password

	if err := s.db.Save(&existing).Error; err != nil {
		return err
	}

	return nil
}

func (s *userService) DeleteUser(id string) error {
	var user models.User
	if err := s.db.First(&user, id).Error; err != nil {
		return err
	}
	return s.db.Delete(&user).Error
}

func (s *userService) GetCountires() ([]string, error) {
	var countries []string
	var user models.User

	err := s.db.Find(&user).Distinct("country").Pluck("country", &countries).Error
	if err != nil {
		return nil, err
	}

	return countries, nil
}
