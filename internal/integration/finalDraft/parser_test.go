package finaldraft

import (
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	valueobjects "github.com/hernangonzalez1987/scriptBreakdown/internal/domain/valueObjects"
)

const sceneText = `
<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<FinalDraft DocumentType="Script" Template="No" Version="2">
  <Content>
    <Paragraph Type="Scene Heading" Alignment="Left">
      <Text Style="AllCaps">EXT. PRAÇA DA SÉ - DIA</Text>
    </Paragraph>
  </Content>
</FinalDraft> 
`

func Test_extractCategoryTagsFromScript(t *testing.T) {
	t.Parallel()

	type args struct {
		script FDXFile
	}
	tests := []struct {
		name string
		args args
		want []entity.Category
	}{
		{
			name: "should extract categories from script",
			args: args{script: FDXFile{
				TagData: TagData{
					TagCategories: TagCategories{
						TagCategories: []TagCategory{
							{
								ID:     "SomeID1",
								Name:   "Cast Members",
								Number: "1",
							},
							{
								ID:     "SomeID2",
								Name:   "Background Actors",
								Number: "2",
							},
						},
					},
				},
			}},
			want: []entity.Category{
				{
					ID:     "SomeID1",
					Type:   valueobjects.TagCategoryCastMembers,
					Number: 1,
				},
				{
					ID:     "SomeID2",
					Type:   valueobjects.TagCategoryBackgroundActors,
					Number: 2,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := extractCategoryTagsFromScript(tt.args.script); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractCategoryTagsFromScript() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extractScenesFromScript_should_join_scene_text(t *testing.T) {
	t.Parallel()

	type args struct {
		script FDXFile
	}
	tests := []struct {
		name string
		args args
		want []entity.Scene
	}{
		{
			name: "should join all the text of a single scene",
			args: args{
				script: FDXFile{
					Content: Content{
						Paragraph: []Paragraph{
							{
								Type:   sceneHeading,
								Number: 2,
							},
							{
								Type: actionHeading,
								Text: []Text{
									{
										Value: "Some scene text.",
									},
									{
										Value: "Some more scene text.",
									},
								},
							},
							{
								Type: actionHeading,
								Text: []Text{
									{
										Value: "Some additional text.",
									},
								},
							},
						},
					},
				},
			},
			want: []entity.Scene{
				{
					Number: 2,
					Text:   " Some scene text.\n Some more scene text.\n Some additional text.\n",
				},
			},
		},
		{
			name: "should number a non numbered scene",
			args: args{
				script: FDXFile{
					Content: Content{
						Paragraph: []Paragraph{
							{
								Type: sceneHeading,
							},
							{
								Type: actionHeading,
								Text: []Text{
									{
										Value: "Some scene text.",
									},
								},
							},
						},
					},
				},
			},
			want: []entity.Scene{
				{
					Number: 1,
					Text:   " Some scene text.\n",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := extractScenesFromScript(tt.args.script); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractScenesFromScript() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readFile(t *testing.T) {
	t.Parallel()

	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name     string
		args     args
		wantFile *FDXFile
		wantHash string
		wantErr  bool
	}{
		{
			name: "should read file",
			args: args{strings.NewReader(sceneText)},
			wantFile: &FDXFile{Content: Content{Paragraph: []Paragraph{
				{
					Number: 0,
					Type:   sceneHeading,
					Text: []Text{{
						Value: "EXT. PRAÇA DA SÉ - DIA",
					}},
				},
			}}},
			wantHash: "58d6af83-6160-3c89-8e37-5ebfb1872392",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gotFile, gotHash, err := readScript(tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("readFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFile, tt.wantFile) {
				t.Errorf("readFile() gotFile = %v, want %v", gotFile, tt.wantFile)
			}
			if gotHash != tt.wantHash {
				t.Errorf("readFile() gotHash = %v, want %v", gotHash, tt.wantHash)
			}
		})
	}
}
