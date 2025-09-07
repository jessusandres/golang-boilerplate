package utils

import (
	"net/http/httptest"
	"testing"

	"lookerdevelopers/boilerplate/cmd/types"

	"github.com/gin-gonic/gin"
)

func TestExtractStateFailForType(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Set("state", "test")

	_, ok := ExtractAppState(c)

	if ok {
		t.Errorf("ExtractAppState should return false")
	}
}

func TestExtractStateSuccess(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Set("state", types.AppState{
		Uuid: "test",
	})

	_, ok := ExtractAppState(c)

	if !ok {
		t.Errorf("ExtractAppState should return true")
	}
}
