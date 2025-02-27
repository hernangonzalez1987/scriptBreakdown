package finaldraft

import (
	"context"
	"encoding/xml"
	"strings"
	"testing"

	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_renderScript(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	source := strings.NewReader(
		`<Paragraph Type="Scene Heading" Number="1"></Paragraph>` +
			`<Paragraph Type="Action"><Text>Some Text </Text></Paragraph><SomeOtherField></SomeOtherField>`)

	target := new(strings.Builder)

	breakdown := entity.ScriptBreakdown{
		SceneBreakdowns: []entity.SceneBreakdown{
			{
				Number: 1,
				Tags: []entity.Tag{
					{
						Number: 1,
						Label:  "Some",
					},
				},
			},
		},
	}

	err := (&Render{}).RenderScript(ctx, source, target, breakdown)

	expected := `<Paragraph Type="Scene Heading" Number="1"></Paragraph>` +
		`<Paragraph Type="Action"><Text TagNumber="1">Some</Text><Text> Text </Text>` +
		`<UserParagraphData></UserParagraphData></Paragraph><SomeOtherField></SomeOtherField>`

	require.NoError(t, err)
	assert.Equal(t, expected, target.String())
}

func Test_processText(t *testing.T) {
	t.Parallel()

	type args struct {
		text           Text
		sceneBreakdown entity.SceneBreakdown
	}

	tests := []struct {
		name string
		args args
		want []Text
	}{
		{
			name: "should not find any tag",
			args: args{
				text:           Text{Value: "some text with no tags"},
				sceneBreakdown: entity.SceneBreakdown{Tags: []entity.Tag{{Label: "not existing tag"}}},
			},
			want: []Text{{Value: "some text with no tags"}},
		},
		{
			name: "should find tag at the begginning of the text",
			args: args{
				text:           Text{Value: "some text with tags"},
				sceneBreakdown: entity.SceneBreakdown{Tags: []entity.Tag{{Label: "some text", Number: 1}}},
			},
			want: []Text{{Value: "some text", TagNumber: "1"}, {Value: " with tags"}},
		},
		{
			name: "should find tag at the end of the text",
			args: args{
				text:           Text{Value: "some text with tags"},
				sceneBreakdown: entity.SceneBreakdown{Tags: []entity.Tag{{Label: "with tags", Number: 2}}},
			},
			want: []Text{{Value: "some text "}, {Value: "with tags", TagNumber: "2"}},
		},
		{
			name: "should find tag at the middle of the text",
			args: args{
				text:           Text{Value: "some text with tags"},
				sceneBreakdown: entity.SceneBreakdown{Tags: []entity.Tag{{Label: "with", Number: 3}}},
			},
			want: []Text{{Value: "some text "}, {Value: "with", TagNumber: "3"}, {Value: " tags"}},
		},
		{
			name: "should find several tags",
			args: args{
				text:           Text{Value: "some text with tags"},
				sceneBreakdown: entity.SceneBreakdown{Tags: []entity.Tag{{Label: "some", Number: 4}, {Label: "tags", Number: 5}}},
			},
			want: []Text{{Value: "some", TagNumber: "4"}, {Value: " text with "}, {Value: "tags", TagNumber: "5"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, tagText(tt.args.text, tt.args.sceneBreakdown))
		})
	}
}

func Test_processTagData(t *testing.T) {
	t.Parallel()

	source := strings.NewReader(
		`<TagData><TagCategories></TagCategories></TagData>`)

	target := new(strings.Builder)

	decoder := xml.NewDecoder(source)
	encoder := xml.NewEncoder(target)

	breakdown := entity.ScriptBreakdown{
		SceneBreakdowns: []entity.SceneBreakdown{
			{
				Number: 1,
				Tags: []entity.Tag{
					{
						ID:     "someTagID",
						Number: 1,
						Label:  "Some",
						Category: entity.Category{
							Number: 1,
							ID:     "someCatID",
						},
					},
				},
			},
		},
	}

	token, _ := decoder.Token()

	var err error

	v, _ := token.(xml.StartElement)
	err = processTagData(&v, decoder, encoder, breakdown)

	expected := `<TagData><TagCategories></TagCategories><TagDefinitions><TagDefinition CatId="someCatID"` +
		` Id="someTagID" Label="Some" Number="1">` +
		`</TagDefinition></TagDefinitions><Tags><Tag Number="1"><DefId>someTagID</DefId></Tag></Tags></TagData>`

	require.NoError(t, err)
	assert.Equal(t, expected, target.String())
}
