package main

import (
	"net/http"

	"projects/pyp-api/internal"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", getPaper)

	router.Run(":8080")
}

func getPaper(c *gin.Context) {
	p, err := internal.ParseCode(c.Query("code"), c.Query("type"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": 0,
			"message": err.Error(),
		})
		return
	}

	url, err := internal.GetPaperUrl(p)
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
