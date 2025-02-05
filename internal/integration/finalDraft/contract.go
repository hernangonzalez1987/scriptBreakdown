package finaldraft

import "encoding/xml"

const (
	SceneHeading  = "Scene Heading"
	ActionHeading = "Action"
)

type FDXFile struct {
	FinalDraft xml.Name `xml:"FinalDraft"`
	Content    Content  `xml:"Content"`
}

type Content struct {
	Content   xml.Name    `xml:"Content"`
	Paragraph []Paragraph `xml:"Paragraph"`
}

type Paragraph struct {
	Paragraph xml.Name `xml:"Paragraph"`
	Number    int      `xml:"Number,attr"`
	Type      string   `xml:"Type,attr"`
	Text      string   `xml:"Text"`
}
