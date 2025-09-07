package utils

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockContext struct {
	*gin.Context
	AbortCalled bool
}

func (m *MockContext) Abort() {
	m.AbortCalled = true
	return
}

func TestHandleServiceErrorWhenErrorIsNil(t *testing.T) {
	recorder := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(recorder)
	mockCtx := &MockContext{
		Context:     c,
		AbortCalled: false,
	}

	hasErr := GinAbortError(c, nil)

	assert.False(t, mockCtx.AbortCalled, "Expected c.Abort() to be called")

	if hasErr {
		t.Errorf("GinAbortError should return false")
	}
}

func TestHandleServiceErrorWhenErrorIsNotNil(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	mockCtx := &MockContext{
		Context:     c,
		AbortCalled: false,
	}

	hasErr := GinAbortError(mockCtx, errors.New("test"))

	assert.True(t, mockCtx.AbortCalled, "Expected c.Abort() to be called")

	if !hasErr {
		t.Errorf("GinAbortError should return true")
	}
}
