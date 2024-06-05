// controllers/country.go
package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// FetchCountries fetches countries from an external API and stores them in the database
// @Summary Fetch countries from external API
// @Description Fetch countries from an external API and store them in the database
// @Tags country
// @Security BearerAuth
// @Produce json
// @Param Authorization header string true "Authorization token"
// @Success 200 {array} models.Country
// @Failure 500 {object} gin.H
// @Router /fetch-countries [get]
func FetchCountries(c *gin.Context) {
	client := &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   20 * time.Second,
			ResponseHeaderTimeout: 20 * time.Second,
			ExpectContinueTimeout: 2 * time.Second,
		},
	}

	resp, err := client.Get("https://api.first.org/data/v1/countries")
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

	var countries interface{}
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
// @Security BearerAuth
// @Produce json
// @Param Authorization header string true "Authorization token"
// @Success 200 {array} models.Country
// @Failure 500 {object} gin.H
// @Router /countries [get]
func (ctrl *UserController) GetCountries(c *gin.Context) {
	countries, err := ctrl.service.GetCountires()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": countries})
}
