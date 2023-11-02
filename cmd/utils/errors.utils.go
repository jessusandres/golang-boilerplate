package utils

import (
	"lookerdevelopers/boilerplate/cmd/interfaces"
)

func HandleServiceError(c interfaces.Aborter, err error) bool {
	if err != nil {
		_ = c.Error(err)
		c.Abort()

		return true
	}

	return false
}
