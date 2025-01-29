package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	usecasebreakdown "github.com/hernangonzalez1987/scriptBreakdown/internal/domain/useCase/breakdown"
	sceneanalyzer "github.com/hernangonzalez1987/scriptBreakdown/internal/domain/useCase/breakdown/sceneAnalyzer"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/integration/deepseek"
	finaldraft "github.com/hernangonzalez1987/scriptBreakdown/internal/integration/finalDraft"
	presentationbreakdown "github.com/hernangonzalez1987/scriptBreakdown/internal/presentation"
)

var deepSeekHost = os.Getenv("DEEPSEEK_HOST")

func main() {

	router := gin.Default()

	router.POST("/script/breakdown", presentationbreakdown.New(
		usecasebreakdown.New(
			finaldraft.New(),
			sceneanalyzer.New(
				deepseek.New(&http.Client{}, deepSeekHost),
			),
		),
	).ProcessFile)

	router.Run(":9000")
}
