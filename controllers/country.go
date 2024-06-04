// controllers/country.go
package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"go-rest-api/database"
	"go-rest-api/models"

	"github.com/gin-gonic/gin"
)

// FetchCountries fetches countries from an external API and stores them in the database
// @Summary Fetch countries from external API
// @Description Fetch countries from an external API and store them in the database
// @Tags country
// @Produce json
// @Success 200 {array} models.Country
// @Failure 500 {object} gin.H
// @Router /countries/fetch [get]
func FetchCountries(c *gin.Context) {
	resp, err := http.Get("https://restcountries.com/v3.1/all")
	if err != nil {
		log.Printf("Error fetching countries: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch countries"})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
		return
	}

	var countries []models.Country
	if err := json.Unmarshal(body, &countries); err != nil {
		log.Printf("Error unmarshaling JSON: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse countries"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": countries})
}

// GetCountries returns all countries stored in the database
// @Summary Get all countries
// @Description Get a list of all countries stored in the database
// @Tags country
// @Produce json
// @Success 200 {array} models.Country
// @Failure 500 {object} gin.H
// @Router /countries [get]
func GetCountries(c *gin.Context) {
	var countries []models.Country
	database.DB.Find(&countries)
	c.JSON(http.StatusOK, gin.H{"data": countries})
}
