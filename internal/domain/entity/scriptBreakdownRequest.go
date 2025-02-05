package entity

import (
	"path/filepath"

	"github.com/go-playground/validator/v10"
)

var validFileExtentions = map[string]bool{".fdx": true}

type ScriptBreakdownRequest struct {
	FilePath string
}

func ScriptBreakdownRequestValidate(sl validator.StructLevel) {

	value, _ := sl.Current().Interface().(ScriptBreakdownRequest)

	fileExtention := filepath.Ext(value.FilePath)

	if validFileExtentions[fileExtention] {
		return
	}

	sl.ReportError(sl.Current().Interface(), "FilePath", "file-path", "custom", fileExtention)

}
