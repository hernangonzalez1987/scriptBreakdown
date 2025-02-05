package entity

type Script struct {
	Hash   string
	Scenes []Scene
}

type Scene struct {
	Number int
	Text   string
}
