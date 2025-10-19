package utils

import (
	"github.com/gin-gonic/gin"
)

// GinAbortError aborts the current request and returns true if the error is not nil, otherwise returns false.
func GinAbortError(c *gin.Context, err error) bool {
	if err != nil {
		_ = c.Error(err)
		c.Abort()

		return true
	}

	return false
}
