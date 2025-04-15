package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   any    `json:"error,omitempty"`
}

func ResOk(c *gin.Context, status int, data any) {
	c.AbortWithStatusJSON(status, Response{
		Status:  status,
		Message: "ok",
		Data:    data,
	})
}

func ResErr(c *gin.Context, status int, message string) {
	c.AbortWithStatusJSON(status, Response{
		Status:  status,
		Message: message,
	})
}

func ResSuccess(c *gin.Context, data any) {
	ResOk(c, http.StatusOK, data)
}

func ResCreated(c *gin.Context, data any) {
	ResOk(c, http.StatusCreated, data)
}

func ResBadRequest(c *gin.Context, message string) {
	ResErr(c, http.StatusBadRequest, message)
}

func ResUnauthorized(c *gin.Context, message string) {
	ResErr(c, http.StatusUnauthorized, message)
}

func ResForbidden(c *gin.Context, message string) {
	ResErr(c, http.StatusForbidden, message)
}

func ResNotFound(c *gin.Context, message string) {
	ResErr(c, http.StatusNotFound, message)
}

func ResInternalServerError(c *gin.Context, message string) {
	ResErr(c, http.StatusInternalServerError, message)
}
