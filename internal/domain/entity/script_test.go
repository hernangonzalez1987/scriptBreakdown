package entity

import (
	"reflect"
	"testing"

	"github.com/go-playground/assert/v2"
	valueobjects "github.com/hernangonzalez1987/scriptBreakdown/internal/domain/valueObjects"
)

func TestTagCategories_GetByType(t *testing.T) {
	t.Parallel()

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
			t.Parallel()

			if got := tt.ref.GetByType(tt.args.ctgType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TagCategories.GetByType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDefaultCategories(t *testing.T) {
	tests := []struct {
		name string
		want []Category
	}{
		{
			name: "should return default categories",
			want: []Category{
				{ID: "ee6c0e4c-b47f-44d9-be5b-e0a87e88b1ee", Type: 9, Number: 0},
				{ID: "2bf0fefc-d647-413b-ad99-a0c37825507d", Type: 18, Number: 1},
				{ID: "19e49e50-19ea-4dee-af6a-5c6b9f3d5393", Type: 2, Number: 2},
				{ID: "77cf9c24-74a9-4971-bb22-2e1320fa74b4", Type: 17, Number: 3},
				{ID: "b5e60211-21ef-4158-baf1-816202124325", Type: 1, Number: 4},
				{ID: "649a196c-3f2d-4d8c-afcc-ca8441871a8f", Type: 8, Number: 5},
				{ID: "6329e32b-3274-4518-bb32-de7e2740e31d", Type: 16, Number: 6},
				{ID: "6d664552-ed0b-4263-bbfa-e0dbb56ee057", Type: 10, Number: 7},
				{ID: "59974f84-1f44-43a5-bf45-f5ea7b501c50", Type: 5, Number: 8},
				{ID: "689a0f83-5c47-4af7-b89e-c871009ada2b", Type: 14, Number: 9},
				{ID: "757f626c-b8f2-4ca0-a51f-e4e1dd1eac25", Type: 12, Number: 10},
				{ID: "12815601-c0b8-48ec-bbd9-abfa136bf503", Type: 11, Number: 11},
				{ID: "944a03b8-8091-433f-93cd-f1dffa9b4d9c", Type: 6, Number: 12},
				{ID: "ec3b599e-104f-4b5f-9287-3f2e7c7b269e", Type: 13, Number: 13},
				{ID: "36e938c9-1250-49da-b8a8-853a555929aa", Type: 3, Number: 14},
				{ID: "a268df9e-97f5-47da-8fa6-6896700c5c7d", Type: 4, Number: 15},
				{ID: "55e7536b-1f38-438d-9561-bc162c1eaac8", Type: 15, Number: 16},
				{ID: "e5fa69db-c851-45db-9a7c-9d47d2d0484c", Type: 7, Number: 17},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetDefaultCategories()
			assert.Equal(t, got, tt.want)
		})
	}
}
