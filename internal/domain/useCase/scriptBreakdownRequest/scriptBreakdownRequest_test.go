package scriptbreakdownrequest

import (
	"bytes"
	"context"
	"errors"
	"os"
	"reflect"
	"testing"

	"github.com/hernangonzalez1987/scriptBreakdown/internal/_mocks"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	valueobjects "github.com/hernangonzalez1987/scriptBreakdown/internal/domain/valueObjects"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestRequestScriptBreakdown(t *testing.T) {
	mockStorage := _mocks.NewMockStorage(t)
	repository := _mocks.NewMockBreakdownRepository(t)
	useCase := New(mockStorage, repository)
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

func TestScriptBreakdownRequestUseCase_GetResult(t *testing.T) {
	storage := _mocks.NewMockStorage(t)
	repository := _mocks.NewMockBreakdownRepository(t)
	ctx := context.Background()

	nonExistingID := "nonExistingID"
	existingIDProcessing := "existingIDProcessing"
	existingIDSuccess := "existingIDSuccess"

	repository.EXPECT().Get(ctx, nonExistingID).Return(nil, errors.New("non existing on db"))
	repository.EXPECT().Get(ctx, existingIDProcessing).Return(
		&entity.ScriptBreakdownResult{Status: valueobjects.BreakdownStatusProcessing}, nil,
	)
	repository.EXPECT().Get(ctx, existingIDSuccess).Return(
		&entity.ScriptBreakdownResult{Status: valueobjects.BreakdownStatusSuccess}, nil,
	)

	file, _ := os.CreateTemp("", "someFile")
	defer file.Close()

	storage.EXPECT().Get(ctx, existingIDSuccess).Return(file, nil)

	useCase := New(storage, repository)

	type args struct {
		breakdownID string
	}
	tests := []struct {
		name    string
		args    args
		want    *entity.ScriptBreakdownResult
		wantErr bool
	}{
		{
			name:    "should return error if breakdown does not exists",
			args:    args{breakdownID: nonExistingID},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "should return pending breakdown",
			args:    args{breakdownID: existingIDProcessing},
			want:    &entity.ScriptBreakdownResult{Status: valueobjects.BreakdownStatusProcessing},
			wantErr: false,
		},
		{
			name: "should return error if breakdown does not exists",
			args: args{breakdownID: existingIDSuccess},
			want: &entity.ScriptBreakdownResult{
				Status:  valueobjects.BreakdownStatusSuccess,
				Content: file,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := useCase.GetResult(ctx, tt.args.breakdownID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ScriptBreakdownRequestUseCase.GetResult() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScriptBreakdownRequestUseCase.GetResult() = %v, want %v", got, tt.want)
			}
		})
	}

	mock.AssertExpectationsForObjects(t, storage, repository)
}
