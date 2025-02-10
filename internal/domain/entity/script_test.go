package entity

import (
	"reflect"
	"testing"

	valueobjects "github.com/hernangonzalez1987/scriptBreakdown/internal/domain/valueObjects"
)

func TestTagCategories_GetByType(t *testing.T) {
	type args struct {
		ctgType valueobjects.TagCategoryType
	}
	tests := []struct {
		name string
		ref  TagCategories
		args args
		want *Category
	}{
		{
			name: "should return category",
			ref: TagCategories{
				{Type: valueobjects.TagCategoryProps},
				{Type: valueobjects.TagCategoryCastMembers},
			},
			args: args{ctgType: valueobjects.TagCategoryProps},
			want: &Category{Type: valueobjects.TagCategoryProps},
		},
		{
			name: "should not return missing category",
			ref: TagCategories{
				{Type: valueobjects.TagCategoryProps},
				{Type: valueobjects.TagCategoryCastMembers},
			},
			args: args{ctgType: valueobjects.TagCategoryUnknown},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ref.GetByType(tt.args.ctgType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TagCategories.GetByType() = %v, want %v", got, tt.want)
			}
		})
	}
}
