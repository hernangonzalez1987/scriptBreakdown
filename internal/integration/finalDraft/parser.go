package finaldraft

import (
	"context"
	"encoding/xml"
	"fmt"
	"os"

	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	"github.com/pkg/errors"
)

type Parser struct{}

func New() *Parser {
	return &Parser{}
}

func (f *Parser) ParseScript(_ context.Context,
	breakdownRequest entity.ScriptBreakdownRequest,
) (*entity.Script, error) {
	fileRawContent, err := os.ReadFile(breakdownRequest.FilePath)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var file FDXFile

	err = xml.Unmarshal(fileRawContent, &file)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var script entity.Script

	var scene entity.Scene

	for _, paragraph := range file.Content.Paragraph {
		if paragraph.Type == "Scene Heading" {
			if scene.Text != "" {
				script.Scenes = append(script.Scenes, scene)
			}

			scene = entity.Scene{
				Number: paragraph.Number,
				Text:   paragraph.Text,
			}

			continue
		}

		if paragraph.Type == "Action" {
			scene.Text = fmt.Sprintln(scene.Text, paragraph.Text)
		}
	}

	if scene.Text != "" {
		script.Scenes = append(script.Scenes, scene)
	}

	return &script, nil
}
