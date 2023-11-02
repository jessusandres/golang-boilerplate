package interfaces

import "github.com/gin-gonic/gin"

type Aborter interface {
	Abort()
	Error(err error) *gin.Error
}
