package valueobjects

import (
	"maps"
	"slices"
)

type TagCategoryType int

const (
	TagCategoryUnknown = iota
	TagCategoryCastMembers
	TagCategoryBackgroundActors
	TagCategoryStunts
	TagCategoryVehicles
	TagCategoryProps
	TagCategorySpecialEffects
	TagCategoryWardrobe
	TagCategoryMakeupHair
	TagCategoryAnimals
	TagCategoryMusic
	TagCategorySound
	TagCategorySetDressing
	TagCategorySpecialEquipment
	TagCategorySecurity
	TagCategoryVisualEffects
	TagCategoryMechanicalEffects
	TagCategoryCamera
	TagCategoryArtDepartment
)

func getNamesMap() map[TagCategoryType]string {
	return map[TagCategoryType]string{
		TagCategoryCastMembers:       "cast",
		TagCategoryBackgroundActors:  "background actors",
		TagCategoryStunts:            "stunts",
		TagCategoryVehicles:          "vehicles",
		TagCategoryProps:             "props",
		TagCategorySpecialEffects:    "special effects",
		TagCategoryWardrobe:          "wardrove",
		TagCategoryMakeupHair:        "makeup/hair",
		TagCategoryAnimals:           "animals",
		TagCategoryMusic:             "music",
		TagCategorySound:             "sound",
		TagCategorySetDressing:       "set dressing",
		TagCategorySpecialEquipment:  "special equipment",
		TagCategorySecurity:          "security",
		TagCategoryVisualEffects:     "visual effects",
		TagCategoryMechanicalEffects: "mechanical effects",
		TagCategoryCamera:            "camera",
		TagCategoryArtDepartment:     "art department",
	}
}

func (tag TagCategoryType) String() string {
	return getNamesMap()[tag]
}

func GetAllTagCategoryNames() []string {
	categoryNames := make([]string, 0)

	for name := range maps.Values(getNamesMap()) {
		categoryNames = append(categoryNames, name)
	}

	slices.Sort[[]string](categoryNames)

	return categoryNames
}

func GetTagCategoryByName(name string) TagCategoryType {
	for tagCategoryType, tagCategoryName := range getNamesMap() {
		if name == tagCategoryName {
			return tagCategoryType
		}
	}

	return TagCategoryUnknown
}
