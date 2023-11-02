package utils

import (
	"cmpc.cl/biopack-cl-boxboard-cf-tracking-provider-order/cmd/types"

	"github.com/gin-gonic/gin"
)

func ExtractState(c *gin.Context) (types.AppState, bool) {
	reqState, exists := c.Get("state")

	if !exists {
		return types.AppState{}, false
	}

	state, ok := reqState.(types.AppState)

	return state, ok
}
