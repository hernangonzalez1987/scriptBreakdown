package csv

import (
	"context"
	"encoding/csv"
	"io"
	"strings"

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
		intExt, location, dayNight := extractLocationDayNight(sceneBreakdown.Header)

		for _, tag := range sceneBreakdown.Tags {
			records = append(records, CsvRecord{
				SceneNumber:   sceneBreakdown.Number,
				SceneIntExt:   intExt,
				SceneLocation: location,
				SceneDayNight: dayNight,
				TagCategory:   tag.Category.Type.String(),
				TagLabel:      tag.Label,
			})
		}
	}

	err := gocsv.MarshalCSV(records, writer)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func extractLocationDayNight(header string) (string, string, string) {
	parts1 := strings.Split(header, ".")

	if len(parts1) <= 1 {
		return header, "", ""
	}

	intExt := parts1[0]

	parts2 := strings.Split(parts1[1], "-")
	if len(parts2) <= 1 {
		return intExt, parts2[0], ""
	}

	return strings.TrimSpace(intExt),
		strings.TrimSpace(parts2[0]),
		strings.TrimSpace(parts2[len(parts2)-1])
}
