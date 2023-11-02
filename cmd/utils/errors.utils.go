package utils

import (
	"cmpc.cl/biopack-cl-boxboard-cf-tracking-provider-order/cmd/interfaces"
)

func HandleServiceError(c interfaces.Aborter, err error) bool {
	if err != nil {
		_ = c.Error(err)
		c.Abort()

		return true
	}

	return false
}
