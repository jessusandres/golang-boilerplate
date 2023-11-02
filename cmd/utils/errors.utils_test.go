package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
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

	hasErr := HandleServiceError(c, nil)

	assert.False(t, mockCtx.AbortCalled, "Expected c.Abort() to be called")

	if hasErr {
		t.Errorf("HandleServiceError should return false")
	}
}

func TestHandleServiceErrorWhenErrorIsNotNil(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	mockCtx := &MockContext{
		Context:     c,
		AbortCalled: false,
	}

	hasErr := HandleServiceError(mockCtx, errors.New("test"))

	assert.True(t, mockCtx.AbortCalled, "Expected c.Abort() to be called")

	if !hasErr {
		t.Errorf("HandleServiceError should return true")
	}
}
