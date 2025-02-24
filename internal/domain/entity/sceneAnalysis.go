package entity

import "time"

type SceneAnalysis struct {
	SceneID       string              `dynamodbav:"sceneId"   json:"sceneId"`
	SceneElements map[string][]string `dynamodbav:"analysis"  json:"analysis"`
	UpdatedAt     time.Time           `dynamodbav:"updatedAt" json:"updatedAt"`
}
