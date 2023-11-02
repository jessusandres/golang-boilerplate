package utils

import (
	"net/http/httptest"
	"testing"

	"cmpc.cl/biopack-cl-boxboard-cf-tracking-provider-order/cmd/types"
	"github.com/gin-gonic/gin"
)

func TestExtractStateFailForType(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Set("state", "test")

	_, ok := ExtractState(c)

	if ok {
		t.Errorf("ExtractState should return false")
	}
}

func TestExtractStateSuccess(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Set("state", types.AppState{
		Uuid: "test",
	})

	_, ok := ExtractState(c)

	if !ok {
		t.Errorf("ExtractState should return true")
	}
}
