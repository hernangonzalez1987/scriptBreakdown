package csv

type CsvRecord struct {
	SceneNumber   int    `csv:"scene_number"`
	SceneLocation string `csv:"scene_location"`
	SceneDayNight string `csv:"scene_day_night"`
	TagCategory   string `csv:"tag_category"`
	TagLabel      string `csv:"tag_label"`
}
