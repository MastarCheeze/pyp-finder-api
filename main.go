package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", getPaper)

	router.Run()
}

func getPaper(c *gin.Context) {
	p, err := parseCode(c.Query("code"), c.Query("type"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": 0,
			"message": err.Error(),
		})
		return
	}

	url, err := getPaperUrl(p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": 0,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": 1,
		"url":     url,
	})
}
