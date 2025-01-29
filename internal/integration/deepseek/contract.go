package deepseek

import "github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"

type DeepSeekContract struct {
	Props []string `json:"props"`
}

func (ref *DeepSeekContract) ToDomain() []entity.Tag {

	tags := []entity.Tag{}

	for _, prop := range ref.Props {
		tags = append(tags, entity.Tag{
			Category: "prop",
			Name:     prop,
		})
	}

	return tags
}
