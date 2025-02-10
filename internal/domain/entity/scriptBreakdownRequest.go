package entity

import (
	"path/filepath"

	"github.com/go-playground/validator/v10"
)

type ScriptBreakdownRequest struct {
	ScriptFileName string
}

func ScriptBreakdownRequestValidate(structVar validator.StructLevel) {
	validFileExtentions := map[string]bool{".fdx": true}

	value, _ := structVar.Current().Interface().(ScriptBreakdownRequest)

	fileExtention := filepath.Ext(value.ScriptFileName)

	if validFileExtentions[fileExtention] {
		return
	}

	structVar.ReportError(structVar.Current().Interface(), "FilePath", "file-path", "custom", fileExtention)
}
