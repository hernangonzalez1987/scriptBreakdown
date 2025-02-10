package entity

import valueobjects "github.com/hernangonzalez1987/scriptBreakdown/internal/domain/valueObjects"

type ScriptBreakdown struct {
	SceneBreakdowns []SceneBreakdown
}

type SceneBreakdown struct {
	Number int
	Tags   []Tag
}

type Tag struct {
	ID       string
	Number   int
	Category Category
	Label    string
}

type Category struct {
	ID     string
	Type   valueobjects.TagCategoryType
	Number int
}

func (ref *ScriptBreakdown) GetSceneBreakdownByNumber(number int) *SceneBreakdown {
	for _, sb := range ref.SceneBreakdowns {
		if sb.Number == number {
			return &sb
		}
	}

	return nil
}
