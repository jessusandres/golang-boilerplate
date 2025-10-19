package utils

import (
	"lookerdevelopers/boilerplate/internal/shared/types"

	"github.com/gin-gonic/gin"
)

// ExtractAppState extracts the key "state" from the context and returns it if it exists, otherwise returns
// an empty state with a false boolean.
func ExtractAppState(c *gin.Context) (types.AppState, bool) {
	reqState, exists := c.Get("state")

	if !exists {
		return types.AppState{}, false
	}

	state, ok := reqState.(types.AppState)

	return state, ok
}
