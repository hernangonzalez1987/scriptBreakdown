package entity

import "time"

type SceneAnalysis struct {
	SceneID       string              `json:"scene_id" dynamodbav:"scene_id"`
	SceneElements map[string][]string `json:"analysis" dynamodbav:"analysis"`
	UpdatedAt     time.Time           `json:"udpated_at" dynamodbav:"udpated_at"`
}
