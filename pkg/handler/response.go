package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewError(message string) *ErrorResponse {
	return &ErrorResponse{Message: message}
}

func newErrorResponse(c *gin.Context, statuscode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statuscode, NewError(message))
}
