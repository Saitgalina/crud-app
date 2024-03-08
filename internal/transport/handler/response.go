package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Println(fmt.Sprintf("ERROR: %s", message))
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
