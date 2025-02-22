package finaldraft

import (
	"encoding/xml"
	"strconv"

	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	valueobjects "github.com/hernangonzalez1987/scriptBreakdown/internal/domain/valueObjects"
)

const (
	sceneHeading  = "Scene Heading"
	actionHeading = "Action"
	paragraphType = "Paragraph"
	typeName      = "Type"
	numberAttr    = "Number"
	tagData       = "TagData"
)

type FDXFile struct {
	FinalDraft xml.Name `xml:"FinalDraft"`
	Content    Content  `xml:"Content"`
	TagData    TagData  `xml:"TagData"`
}

type Content struct {
	Content   xml.Name    `xml:"Content"`
	Paragraph []Paragraph `xml:"Paragraph"`
}

type Paragraph struct {
	Type              string            `xml:"Type,attr,omitempty"`
	Number            int               `xml:"Number,attr,omitempty"`
	Text              []Text            `xml:"Text,omitempty"`
	UserParagraphData UserParagraphData `xml:"UserParagraphData,omitempty"`
}

type Text struct {
	Value     string `xml:",innerxml"`
	TagNumber string `xml:"TagNumber,attr,omitempty"`
}

type UserParagraphData struct {
	Value string `xml:",innerxml"`
}

type TagData struct {
	TagCategories  TagCategories  `xml:"TagCategories"`
	TagDefinitions TagDefinitions `xml:"TagDefinitions"`
	Tags           Tags           `xml:"Tags"`
}

type TagCategories struct {
	XMLName       xml.Name      `xml:"TagCategories"`
	TagCategories []TagCategory `xml:"TagCategory"`
}

type TagDefinitions struct {
	XMLName        xml.Name        `xml:"TagDefinitions"`
	TagDefinitions []TagDefinition `xml:"TagDefinition"`
}

type Tags struct {
	XMLName xml.Name `xml:"Tags"`
	Tags    []Tag    `xml:"Tag"`
}

type TagCategory struct {
	TagCategory xml.Name `xml:"TagCategory"`
	Name        string   `xml:"Name,attr"`
	Number      string   `xml:"Number,attr"`
	ID          string   `xml:"Id,attr"`
	Style       string   `xml:"Style,attr"`
}

type TagDefinition struct {
	CatId  string `xml:"CatId,attr,omitempty"`
	Id     string `xml:"Id,attr,omitempty"`
	Label  string `xml:"Label,attr,omitempty"`
	Number string `xml:"Number,attr,omitempty"`
}

type Tag struct {
	Number string `xml:"Number,attr,omitempty"`
	DefId  string `xml:"DefId"`
}

func isSceneHeading(token xml.StartElement) bool {
	if token.Name.Local == paragraphType {
		for _, attr := range token.Attr {
			if attr.Name.Local == typeName {
				return attr.Value == sceneHeading
			}
		}
	}

	return false
}

func isActionHeading(token xml.StartElement) bool {
	if token.Name.Local == paragraphType {
		for _, attr := range token.Attr {
			if attr.Name.Local == typeName {
				return attr.Value == actionHeading
			}
		}
	}

	return false
}

func isTagDataElement(token xml.StartElement) bool {
	return token.Name.Local == tagData
}

func getSceneNumber(token xml.StartElement) int {
	for _, attr := range token.Attr {
		if attr.Name.Local == numberAttr {
			n, _ := strconv.Atoi(attr.Value)

			return n
		}
	}

	return 0
}

func (ref TagCategory) ToDomain() *entity.Category {
	number, _ := strconv.Atoi(ref.Number)

	tagCategoryType := translateType(ref.Name)

	if tagCategoryType == valueobjects.TagCategoryUnknown {
		return nil
	}

	return &entity.Category{
		ID:     ref.ID,
		Number: number,
		Type:   tagCategoryType,
	}
}

func translateType(name string) valueobjects.TagCategoryType {
	categoryTypes := map[string]valueobjects.TagCategoryType{
		"Cast Members":       valueobjects.TagCategoryCastMembers,
		"Background Actors":  valueobjects.TagCategoryBackgroundActors,
		"Stunts":             valueobjects.TagCategoryStunts,
		"Vehicles":           valueobjects.TagCategoryVehicles,
		"Props":              valueobjects.TagCategoryProps,
		"Special Effects":    valueobjects.TagCategorySpecialEffects,
		"Wardrobe":           valueobjects.TagCategoryWardrobe,
		"Makeup/Hair":        valueobjects.TagCategoryMakeupHair,
		"Animals":            valueobjects.TagCategoryAnimals,
		"Music":              valueobjects.TagCategoryMusic,
		"Sound":              valueobjects.TagCategorySound,
		"Set Dressing":       valueobjects.TagCategorySetDressing,
		"Special Equipment":  valueobjects.TagCategorySpecialEquipment,
		"Security":           valueobjects.TagCategorySecurity,
		"Visual Effects":     valueobjects.TagCategoryVisualEffects,
		"Mechanical Effects": valueobjects.TagCategoryMechanicalEffects,
		"Camera":             valueobjects.TagCategoryCamera,
		"Art Department":     valueobjects.TagCategoryArtDepartment,
	}

	ctg, found := categoryTypes[name]
	if !found {
		return valueobjects.TagCategoryUnknown
	}

	return ctg
}
