package utils

import (
	"lookerdevelopers/boilerplate/cmd/types"

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
