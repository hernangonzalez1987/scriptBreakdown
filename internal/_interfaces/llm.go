package _interfaces

import "context"

type SceneTextAnalyzer interface {
	AnalyzeSceneText(ctx context.Context, sceneText string) (map[string][]string, error)
}
