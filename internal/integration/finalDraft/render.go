package finaldraft

import (
	"context"
	"encoding/xml"
	"io"
	"strconv"
	"strings"

	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	"github.com/pkg/errors"
)

type Render struct{}

func NewRender() *Render {
	return &Render{}
}

func (ref *Render) RenderScript(_ context.Context, source io.Reader,
	target io.Writer, breakdown entity.ScriptBreakdown,
) error {
	var token xml.Token

	var sceneNumber int

	var sceneCount int

	decoder := xml.NewDecoder(source)
	encoder := xml.NewEncoder(target)

	defer encoder.Close()

	for {
		err := encode(token, encoder)
		if err != nil {
			return errors.WithStack(err)
		}

		token, err = decoder.Token()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return errors.WithStack(err)
		}

		elem, isOk := token.(xml.StartElement)
		if isOk && isSceneHeading(elem) {
			sceneCount++

			sceneNumber = defineSceneNumber(elem, sceneCount)

			continue
		}

		if isOk && isActionHeading(elem) {
			err = processActionParagraph(&elem, sceneNumber, decoder, encoder, breakdown)
			if err != nil {
				return errors.WithStack(err)
			}

			token = nil
		}

		if isOk && isTagDataElement(elem) {
			err = processTagData(&elem, decoder, encoder, breakdown)
			if err != nil {
				return errors.WithStack(err)
			}

			token = nil
		}
	}

	return nil
}

func encode(token xml.Token, encoder *xml.Encoder) error {
	if token != nil {
		return errors.WithStack(encoder.EncodeToken(token))
	}

	return nil
}

func defineSceneNumber(v xml.StartElement, sceneCount int) int {
	sceneNumber := getSceneNumber(v)

	if sceneNumber == 0 {
		sceneNumber = sceneCount
	}

	return sceneNumber
}

func processActionParagraph(token *xml.StartElement, sceneNumber int, decoder *xml.Decoder,
	encoder *xml.Encoder,
	breakdown entity.ScriptBreakdown,
) error {
	sceneBreakdown := breakdown.GetSceneBreakdownByNumber(sceneNumber)
	if sceneBreakdown == nil {
		return nil
	}

	paragraph := Paragraph{}

	err := decoder.DecodeElement(&paragraph, token)
	if err != nil {
		return errors.WithStack(err)
	}

	taggedSceneText := make([]Text, 0)

	for _, text := range paragraph.Text {
		taggedSceneText = append(taggedSceneText, tagText(text, *sceneBreakdown)...)
	}

	paragraph.Text = taggedSceneText
	// should not set Type to avoid duplicating the attr
	paragraph.Type = ""

	err = encoder.EncodeElement(paragraph, *token)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func tagText(text Text, sceneBreakdown entity.SceneBreakdown) []Text {
	if len(text.Value) == 0 {
		return nil
	}

	for _, tag := range sceneBreakdown.Tags {
		pos := strings.Index(text.Value, tag.Label)
		if pos >= 0 {
			prevTexts := tagText(Text{Value: text.Value[:pos]}, sceneBreakdown)

			nextTexts := tagText(Text{Value: text.Value[pos+len(tag.Label):]}, sceneBreakdown)

			taggedText := Text{Value: tag.Label, TagNumber: strconv.Itoa(tag.Number)}

			return append(append(prevTexts, taggedText), nextTexts...)
		}
	}

	return []Text{text}
}

func processTagData(token *xml.StartElement, decoder *xml.Decoder, encoder *xml.Encoder,
	breakdown entity.ScriptBreakdown,
) error {
	tagData := TagData{}

	err := decoder.DecodeElement(&tagData, token)
	if err != nil {
		return errors.WithStack(err)
	}

	for _, sceneBreakdown := range breakdown.SceneBreakdowns {
		for _, tag := range sceneBreakdown.Tags {
			tagData.TagDefinitions.TagDefinitions = append(tagData.TagDefinitions.TagDefinitions,
				TagDefinition{
					CatID:  tag.Category.ID,
					ID:     tag.ID,
					Label:  tag.Label,
					Number: strconv.Itoa(tag.Number),
				})
			tagData.Tags.Tags = append(tagData.Tags.Tags,
				Tag{
					Number: strconv.Itoa(tag.Number),
					DefID:  tag.ID,
				})
		}
	}

	err = encoder.EncodeElement(tagData, *token)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
