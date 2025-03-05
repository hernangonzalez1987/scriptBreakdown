package entity

import valueobjects "github.com/hernangonzalez1987/scriptBreakdown/internal/domain/valueObjects"

type Script struct {
	Scenes        []Scene
	TagCategories TagCategories
}

type TagCategories []Category

type Scene struct {
	Number int
	Header string
	Text   string
}

func (ref TagCategories) GetByType(ctgType valueobjects.TagCategoryType) *Category {
	for _, category := range ref {
		if category.Type == ctgType {
			return &category
		}
	}

	return nil
}
