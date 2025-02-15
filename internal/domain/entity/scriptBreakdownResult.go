package entity

import (
	"io"
	"time"

	valueobjects "github.com/hernangonzalez1987/scriptBreakdown/internal/domain/valueObjects"
)

type ScriptBreakdownResult struct {
	BreakdownID       string                       `json:"BreakdownID"`
	Content           io.Reader                    `json:"Content"`
	Status            valueobjects.BreakdownStatus `json:"Status"`
	StatusDescription string                       `json:"StatusDescription"`
	Version           int                          `json:"Version"`
	LastUpdate        time.Time                    `json:"LastUpdate"`
}
