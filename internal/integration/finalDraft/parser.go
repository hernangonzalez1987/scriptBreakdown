package finaldraft

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"

	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	"github.com/pkg/errors"
)

const (
	namespace = "1486e36b-d00c-4f29-9098-10eb8eab9002"
	sufix     = "_tagged"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (ref *Parser) ParseScript(_ context.Context, reader io.Reader) (*entity.Script, error) {
	file, err := readScript(reader)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &entity.Script{
		Scenes:        extractScenesFromScript(*file),
		TagCategories: extractCategoryTagsFromScript(*file),
	}, nil
}

func readScript(reader io.Reader) (fdxFile *FDXFile, err error) {
	scriptRawContent, err := io.ReadAll(reader)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	fdxFile = &FDXFile{}

	err = xml.Unmarshal(scriptRawContent, fdxFile)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return fdxFile, nil
}

func extractScenesFromScript(script FDXFile) []entity.Scene {
	var scenes []entity.Scene
	var scene entity.Scene
	var sceneCount int

	for _, paragraph := range script.Content.Paragraph {
		if paragraph.Type == sceneHeading {
			if scene.Text != "" {
				scenes = append(scenes, scene)
			}

			sceneCount++

			if paragraph.Number == 0 {
				paragraph.Number = sceneCount
			}

			scene.Number = paragraph.Number
			scene.Text = ""

			continue
		}

		if paragraph.Type == actionHeading {
			for _, text := range paragraph.Text {
				scene.Text = fmt.Sprintln(scene.Text, text.Value)
			}
		}
	}

	if scene.Text != "" {
		scenes = append(scenes, scene)
	}

	return scenes
}

func extractCategoryTagsFromScript(script FDXFile) []entity.Category {
	var tagCategories []entity.Category

	for _, tagCategory := range script.TagData.TagCategories.TagCategories {

		cat := tagCategory.ToDomain()

		if cat != nil {
			tagCategories = append(tagCategories, *cat)
		}
	}

	return tagCategories
}
