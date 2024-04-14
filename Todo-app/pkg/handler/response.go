package handler

import (
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, status int, message string) {
	color.Magenta("\t\t NEW ERROR RESPONSE // from handler ")
	logrus.Errorf(message)
	c.AbortWithStatusJSON(status, errorResponse{message})
}
