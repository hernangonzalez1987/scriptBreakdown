package scriptbreakdownrequest

import (
	"bytes"
	"context"
	"testing"

	"github.com/hernangonzalez1987/scriptBreakdown/internal/_mocks"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestRequestScriptBreakdown(t *testing.T) {
	mockStorage := _mocks.NewMockStorage(t)
	useCase := New(mockStorage)
	ctx := context.Background()

	tempFile := bytes.NewBuffer([]byte("test script content"))

	req := entity.ScriptBreakdownRequest{TempScriptFile: tempFile}

	mockStorage.EXPECT().Put(ctx, "a304328a-1455-35a5-a996-6ea6289a980a",
		mock.AnythingOfType("*os.File")).Return(nil)

	result, err := useCase.RequestScriptBreakdown(ctx, req)

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, &entity.ScriptBreakdownResult{BreakdownID: "a304328a-1455-35a5-a996-6ea6289a980a"}, result)

	mockStorage.AssertExpectations(t)
}
