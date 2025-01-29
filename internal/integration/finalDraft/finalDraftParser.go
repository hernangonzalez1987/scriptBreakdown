package finaldraft

import (
	"context"
	"encoding/xml"
	"fmt"
	"os"

	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/_interfaces"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
)

type finalDraftParser struct {
}

func New() _interfaces.ScriptParser {
	return &finalDraftParser{}
}

func (f *finalDraftParser) ParseScript(ctx context.Context,
	breakdownRequest entity.ScriptBreakdownRequest) (*entity.Script, error) {

	fileRawContent, err := os.ReadFile(breakdownRequest.FilePath)
	if err != nil {
		return nil, err
	}

	file := FDXFile{}

	err = xml.Unmarshal(fileRawContent, &file)
	if err != nil {
		return nil, err
	}

	script := entity.Script{}

	scene := entity.Scene{}

	for _, p := range file.Content.Paragraph {
		if p.Type == "Scene Heading" {
			if scene.Text != "" {
				script.Scenes = append(script.Scenes, scene)
			}
			scene = entity.Scene{
				Number: p.Number,
				Text:   p.Text,
			}
			continue
		}

		if p.Type == "Action" {
			scene.Text = fmt.Sprintln(scene.Text, p.Text)
		}
	}

	return &script, nil

}
