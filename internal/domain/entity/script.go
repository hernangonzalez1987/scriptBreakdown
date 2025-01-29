package entity

type Script struct {
	Scenes []Scene
}

type Scene struct {
	Number int
	Text   string
}
