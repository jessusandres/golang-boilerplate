package incidentsinterfaces

import "github.com/gin-gonic/gin"

type IIncidentsController interface {
	Patch(ctx *gin.Context)
	Get(ctx *gin.Context)
}
