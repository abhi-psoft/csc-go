package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func countryListHandler(c *gin.Context) {
	c.JSON(http.StatusOK, Countries)
}

func stateListHandler(c *gin.Context) {
	countryCode := c.Param("cid")
	var filteredStates = []State{}
	for _, state := range States {
		if state.CountryCode == countryCode {
			filteredStates = append(filteredStates, state)
		}
	}
	c.JSON(http.StatusOK, filteredStates)
}

func cityListHandler(c *gin.Context) {
	countryCode := c.Param("cid")
	stateCode := c.Param("sid")
	var filteredCities = []City{}
	for _, city := range Cities {
		if city.StateCode == stateCode && city.CountryCode == countryCode {
			filteredCities = append(filteredCities, city)
		}
	}
	c.JSON(http.StatusOK, filteredCities)
}

func main() {
	r := gin.Default()

	v1AtxServerList := r.Group("/v1/airattix/server/list")
	{
		v1AtxServerList.GET("/countries", countryListHandler)
		v1AtxServerList.GET("/country/:cid/states", stateListHandler)
		v1AtxServerList.GET("/country/:cid/state/:sid/cities", cityListHandler)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 [default]
}
