package entity

import (
	"io"
)

type ScriptBreakdownRequest struct {
	TempScriptFile io.Reader
}
