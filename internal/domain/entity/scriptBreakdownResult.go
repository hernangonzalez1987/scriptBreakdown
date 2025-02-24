package entity

import (
	"io"
	"time"

	valueobjects "github.com/hernangonzalez1987/scriptBreakdown/internal/domain/valueObjects"
)

type ScriptBreakdownResult struct {
	BreakdownID       string                       `dynamodbav:"breakdownId"       json:"breakdownId"`
	Content           io.ReadCloser                `dynamodbav:"-"                 json:"content"`
	Status            valueobjects.BreakdownStatus `dynamodbav:"status"            json:"status"`
	StatusDescription string                       `dynamodbav:"statusDescription" json:"statusDescription"`
	Version           int                          `dynamodbav:"version"           json:"version"`
	CreatedAt         time.Time                    `dynamodbav:"createAt"          json:"createAt"`
	UpdatedAt         time.Time                    `dynamodbav:"udpatedAt"         json:"udpatedAt"`
}
