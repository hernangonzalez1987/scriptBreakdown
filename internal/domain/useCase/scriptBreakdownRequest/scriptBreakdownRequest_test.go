package scriptbreakdownrequest

import (
	"context"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/hernangonzalez1987/scriptBreakdown/internal/_mocks"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	valueobjects "github.com/hernangonzalez1987/scriptBreakdown/internal/domain/valueObjects"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestRequestScriptBreakdown(t *testing.T) {
	t.Parallel()

	scriptsStorage := _mocks.NewMockStorage(t)
	repository := _mocks.NewMockBreakdownRepository(t)
	useCase := New(scriptsStorage, nil, repository)
	ctx := context.Background()

	tempFile := strings.NewReader("test script content")

	req := entity.ScriptBreakdownRequest{TempScriptFile: tempFile}

	scriptsStorage.EXPECT().Put(ctx, "a304328a-1455-35a5-a996-6ea6289a980a",
		mock.AnythingOfType("*os.File")).Return(nil)

	result, err := useCase.RequestScriptBreakdown(ctx, req)

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, &entity.ScriptBreakdownResult{BreakdownID: "a304328a-1455-35a5-a996-6ea6289a980a"}, result)
}

func TestScriptBreakdownRequestUseCase_GetResult(t *testing.T) {
	breakdownsStorage := _mocks.NewMockStorage(t)
	repository := _mocks.NewMockBreakdownRepository(t)
	ctx := context.Background()

	nonExistingID := "nonExistingID"
	existingIDProcessing := "existingIDProcessing"
	existingIDSuccess := "existingIDSuccess"

	repository.EXPECT().Get(ctx, nonExistingID).Return(nil, errNotFound)
	repository.EXPECT().Get(ctx, existingIDProcessing).Return(
		&entity.ScriptBreakdownResult{Status: valueobjects.BreakdownStatusProcessing}, nil,
	)
	repository.EXPECT().Get(ctx, existingIDSuccess).Return(
		&entity.ScriptBreakdownResult{Status: valueobjects.BreakdownStatusSuccess}, nil,
	)

	file, _ := os.CreateTemp(t.TempDir(), "someFile")
	defer file.Close()

	breakdownsStorage.EXPECT().Get(ctx, existingIDSuccess).Return(file, nil)

	useCase := New(nil, breakdownsStorage, repository)

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

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			got, err := useCase.GetResult(ctx, testCase.args.breakdownID)
			if (err != nil) != testCase.wantErr {
				t.Errorf("ScriptBreakdownRequestUseCase.GetResult() error = %v, wantErr %v", err, testCase.wantErr)

				return
			}

			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("ScriptBreakdownRequestUseCase.GetResult() = %v, want %v", got, testCase.want)
			}
		})
	}
}
