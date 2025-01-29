package entity

type ScriptBreakdown struct {
	SceneBreakdowns []SceneBreakdown
}

type SceneBreakdown struct {
	Number int
	Tags   []Tag
}

type Tag struct {
	Category    string
	Name        string
	Description string
}
