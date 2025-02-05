package finaldraft

import (
	"context"
	"encoding/xml"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/_interfaces"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	"github.com/pkg/errors"
)

const namespace = "1486e36b-d00c-4f29-9098-10eb8eab9002"

type Parser struct {
	cache _interfaces.Cache[FDXFile]
}

func New(cache _interfaces.Cache[FDXFile]) *Parser {
	return &Parser{cache: cache}
}

func (ref *Parser) ParseScript(_ context.Context, req entity.ScriptBreakdownRequest) (*entity.Script, error) {
	fileRawContent, err := os.ReadFile(req.FilePath)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var file FDXFile

	err = xml.Unmarshal(fileRawContent, &file)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	scriptHash := uuid.NewMD5(uuid.MustParse(namespace), fileRawContent)

	ref.cache.Save(scriptHash.String(), file)

	var script entity.Script

	var scene entity.Scene

	for _, paragraph := range file.Content.Paragraph {
		if paragraph.Type == SceneHeading {
			continue
		}

		if paragraph.Type == ActionHeading {
			scene.Text = fmt.Sprintln(scene.Text, paragraph.Text)
		}
	}

	if scene.Text != "" {
		script.Scenes = append(script.Scenes, scene)
	}

	script.Hash = scriptHash.String()

	return &script, nil
}
