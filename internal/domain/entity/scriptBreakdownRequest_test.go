package entity_test

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestScriptBreakdownRequestValidate_should_return_error(t *testing.T) {
	t.Parallel()

	validate := validator.New()

	var req entity.ScriptBreakdownRequest

	validate.RegisterStructValidation(entity.ScriptBreakdownRequestValidate, req)

	request := entity.ScriptBreakdownRequest{
		ScriptFileName: "someFileWithWrongFilePath.txt",
	}

	err := validate.Struct(request)

	assert.Error(t, err)
}

func TestScriptBreakdownRequestValidate_should_return_no_error(t *testing.T) {
	t.Parallel()

	var req entity.ScriptBreakdownRequest

	validate := validator.New()
	validate.RegisterStructValidation(entity.ScriptBreakdownRequestValidate, req)

	request := entity.ScriptBreakdownRequest{
		ScriptFileName: "someFileWithWrongFilePath.fdx",
	}

	err := validate.Struct(request)

	assert.NoError(t, err)
}
