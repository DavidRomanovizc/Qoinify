package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var ErrNotFound = errors.New(http.StatusText(http.StatusNotFound))

func NoMethodHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(405, gin.H{
			"message": "Invalid method",
		})
	}
}

func NoRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(404, gin.H{
			"message": "The processing function of the request route was not found",
		})
	}
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, err := range c.Errors {
			switch err.Err {
			case ErrNotFound:
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": ErrNotFound.Error()})
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			}
		}
	}
}
