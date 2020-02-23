package main

import (
	"github.com/gin-gonic/gin"
)

func createError(c *gin.Context) {
	c.Status(500)
	return
	// c.Fail(http.StatusInternalServerError, fmt.Errorf("custom error"))
	// c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("custom error"))
}
