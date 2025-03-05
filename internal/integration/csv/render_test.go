package csv

import (
	"bytes"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	valueobjects "github.com/hernangonzalez1987/scriptBreakdown/internal/domain/valueObjects"
	"github.com/stretchr/testify/require"
)

func TestRender_RenderScript(t *testing.T) {
	render := Render{}

	buffer := &bytes.Buffer{}

	breakdown := entity.ScriptBreakdown{
		SceneBreakdowns: []entity.SceneBreakdown{
			{
				Number: 1,
				Header: "INT. SOME LOCATION/DETAIL - DAY",
				Tags: []entity.Tag{
					{
						Label:    "someProp",
						Category: entity.Category{Type: valueobjects.TagCategoryProps},
					},
					{
						Label:    "someOtherProp",
						Category: entity.Category{Type: valueobjects.TagCategoryProps},
					},
				},
			},
			{
				Number: 2,
				Header: "EXT. SOME LOCATION/DETAIL - NIGTH",
				Tags: []entity.Tag{
					{
						Label:    "someAnimal",
						Category: entity.Category{Type: valueobjects.TagCategoryAnimals},
					},
				},
			},
		},
	}

	err := render.RenderScript(nil, nil, buffer, breakdown)

	require.NoError(t, err)

	expected := "scene_number,scene_int_ext,scene_location,scene_day_night,tag_category,tag_label\n" +
		"1,INT,SOME LOCATION/DETAIL,DAY,props,someProp\n" +
		"1,INT,SOME LOCATION/DETAIL,DAY,props,someOtherProp\n" +
		"2,EXT,SOME LOCATION/DETAIL,NIGTH,animals,someAnimal\n"

	assert.Equal(t, expected, buffer.String())
}
