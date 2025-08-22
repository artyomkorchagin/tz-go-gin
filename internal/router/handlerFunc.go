package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @name HTTPError
// @description Standard error response returned by the API
// @property Code
// @type integer
// @description HTTP status code
// @property Err
// @type error
// @description Human-readable error message
type HTTPErrorSwaggerWrapper HTTPError

// HTTPError represents an error response returned by the API.
// This is used to standardize error responses across the application.
type HTTPError struct {
	// HTTP status code
	Code int `json:"code"`

	// Error message
	Err error `json:"error"`
}

func (e HTTPError) Error() string {
	return e.Err.Error()
}

type handlerFunc func(c *gin.Context) error

func (h *Handler) wrap(fn handlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := fn(c)
		if err != nil {
			if httpErr, ok := err.(HTTPError); ok {
				h.logger.Error("error", zap.Error(err))
				c.JSON(httpErr.Code, gin.H{"error": httpErr.Err.Error()})
			} else {
				h.logger.Error("error", zap.Error(err))
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
		}
	}
}
