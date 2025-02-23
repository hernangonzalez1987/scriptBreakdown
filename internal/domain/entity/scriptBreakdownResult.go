package entity

import (
	"io"
	"time"

	valueobjects "github.com/hernangonzalez1987/scriptBreakdown/internal/domain/valueObjects"
)

type ScriptBreakdownResult struct {
	BreakdownID       string                       `json:"breakdown_id" dynamodbav:"breakdown_id"`
	Content           io.ReadCloser                `json:"content" dynamodbav:"-"`
	Status            valueobjects.BreakdownStatus `json:"status" dynamodbav:"status"`
	StatusDescription string                       `json:"status_description" dynamodbav:"status_description"`
	Version           int                          `json:"version" dynamodbav:"version"`
	CreatedAt         time.Time                    `json:"created_at" dynamodbav:"created_at"`
	UpdatedAt         time.Time                    `json:"updated_at" dynamodbav:"updated_at"`
}
