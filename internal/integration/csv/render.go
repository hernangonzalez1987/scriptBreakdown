package csv

import (
	"context"
	"encoding/csv"
	"io"

	"github.com/gocarina/gocsv"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	"github.com/pkg/errors"
)

type Render struct{}

func NewRender() *Render {
	return &Render{}
}

func (ref *Render) RenderScript(_ context.Context, source io.Reader,
	target io.Writer, breakdown entity.ScriptBreakdown,
) error {
	writer := gocsv.NewSafeCSVWriter(csv.NewWriter(target))

	records := make([]CsvRecord, 0, len(breakdown.SceneBreakdowns))

	for _, sceneBreakdown := range breakdown.SceneBreakdowns {
		for _, tag := range sceneBreakdown.Tags {
			records = append(records, CsvRecord{
				SceneNumber: sceneBreakdown.Number,
				TagCategory: tag.Category.Type.String(),
				TagLabel:    tag.Label,
			})
		}
	}

	err := gocsv.MarshalCSV(records, writer)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
