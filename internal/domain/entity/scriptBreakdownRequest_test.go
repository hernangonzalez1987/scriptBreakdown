package entity

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestScriptBreakdownRequestValidate_should_return_error(t *testing.T) {

	validate := validator.New()

	validate.RegisterStructValidation(ScriptBreakdownRequestValidate, ScriptBreakdownRequest{})

	request := ScriptBreakdownRequest{
		FilePath: "someFileWithWrongFilePath.txt",
	}

	err := validate.Struct(request)

	assert.Error(t, err)

}

func TestScriptBreakdownRequestValidate_should_return_no_error(t *testing.T) {

	validate := validator.New()

	validate.RegisterStructValidation(ScriptBreakdownRequestValidate, ScriptBreakdownRequest{})

	request := ScriptBreakdownRequest{
		FilePath: "someFileWithWrongFilePath.fdx",
	}

	err := validate.Struct(request)

	assert.NoError(t, err)

}
