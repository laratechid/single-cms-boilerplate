package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int    `json:"status" example:"200"`
	Message string `json:"message" example:"ok"`
	Data    any    `json:"data"`
}

type ResponsePaginate struct {
	Status     int    `json:"status" example:"200"`
	Message    string `json:"message" example:"ok"`
	Data       any    `json:"data"`
	Limit      int    `json:"limit"`
	TotalEntry int    `json:"total_entry"`
}

func ResOk(c *gin.Context, status int, data any) {
	c.AbortWithStatusJSON(status, Response{
		Status:  status,
		Message: "ok",
		Data:    data,
	})
}

func ResPaginate(c *gin.Context, data any, limit int, total int) {
	c.AbortWithStatusJSON(http.StatusOK, ResponsePaginate{
		Status:     http.StatusOK,
		Message:    "ok",
		Data:       data,
		Limit:      limit,
		TotalEntry: total,
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
