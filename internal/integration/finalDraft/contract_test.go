package finaldraft

import (
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	valueobjects "github.com/hernangonzalez1987/scriptBreakdown/internal/domain/valueObjects"
)

var sceneHeadingElement = xml.StartElement{
	Name: xml.Name{
		Space: "",
		Local: "Paragraph",
	},
	Attr: []xml.Attr{
		{
			Name: xml.Name{
				Space: "",
				Local: "Number",
			},
			Value: "1",
		},
		{
			Name: xml.Name{
				Space: "",
				Local: "Type",
			},
			Value: "Scene Heading",
		},
	},
}

var actionHeadingElement = xml.StartElement{
	Name: xml.Name{
		Space: "",
		Local: "Paragraph",
	},
	Attr: []xml.Attr{
		{
			Name: xml.Name{
				Space: "",
				Local: "Type",
			},
			Value: "Action",
		},
	},
}

func Test_isSceneHeading(t *testing.T) {
	t.Parallel()

	type args struct {
		token xml.StartElement
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should identity element as scene heading",
			args: args{token: sceneHeadingElement},
			want: true,
		},
		{
			name: "should not identify element as scene heading",
			args: args{token: actionHeadingElement},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := isSceneHeading(tt.args.token); got != tt.want {
				t.Errorf("isSceneHeading() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isActionHeading(t *testing.T) {
	t.Parallel()

	type args struct {
		token xml.StartElement
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should identity element as action heading",
			args: args{token: actionHeadingElement},
			want: true,
		},
		{
			name: "should not identify element as action heading",
			args: args{token: sceneHeadingElement},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := isActionHeading(tt.args.token); got != tt.want {
				t.Errorf("isActionHeading() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isTagDataElement(t *testing.T) {
	t.Parallel()

	type args struct {
		token xml.StartElement
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should identity element as tag data",
			args: args{token: xml.StartElement{Name: xml.Name{Local: "TagData"}}},
			want: true,
		},
		{
			name: "should not identify element as tag data",
			args: args{token: xml.StartElement{Name: xml.Name{Local: "OtherName"}}},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := isTagDataElement(tt.args.token); got != tt.want {
				t.Errorf("isTextHeading() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSceneNumber(t *testing.T) {
	t.Parallel()

	type args struct {
		token xml.StartElement
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should get valid scene number",
			args: args{token: sceneHeadingElement},
			want: 1,
		},
		{
			name: "should not get valid scene number",
			args: args{token: actionHeadingElement},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := getSceneNumber(tt.args.token); got != tt.want {
				t.Errorf("getSceneNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTagCategory_ToDomain(t *testing.T) {
	t.Parallel()

	type fields struct {
		TagCategory xml.Name
		Name        string
		Number      string
		ID          string
	}
	tests := []struct {
		name   string
		fields fields
		want   *entity.Category
	}{
		{
			name: "should translate category with valid number",
			fields: fields{
				Name:   "Props",
				Number: "1",
				ID:     "someID",
			},
			want: &entity.Category{
				Type:   valueobjects.TagCategoryProps,
				Number: 1,
				ID:     "someID",
			},
		},
		{
			name: "should translate category without number",
			fields: fields{
				Name:   "Props",
				Number: "",
				ID:     "someID",
			},
			want: &entity.Category{
				Type:   valueobjects.TagCategoryProps,
				Number: 0,
				ID:     "someID",
			},
		},
		{
			name: "should not translate category with unknown type",
			fields: fields{
				Name:   "Some Unknown type",
				Number: "1",
				ID:     "someID",
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ref := TagCategory{
				TagCategory: tt.fields.TagCategory,
				Name:        tt.fields.Name,
				Number:      tt.fields.Number,
				ID:          tt.fields.ID,
			}
			if got := ref.ToDomain(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TagCategory.ToDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}
