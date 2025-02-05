package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/useCase/scriptbreakdown"
	scenebreakdown "github.com/hernangonzalez1987/scriptBreakdown/internal/domain/useCase/scriptbreakdown/sceneBreakdown"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/integration/ai"
	finaldraft "github.com/hernangonzalez1987/scriptBreakdown/internal/integration/finalDraft"
	presentationbreakdown "github.com/hernangonzalez1987/scriptBreakdown/internal/presentation"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/tools/cache"
	"github.com/rs/zerolog"
	"github.com/tmc/langchaingo/llms/googleai"
)

func main() {
	logger := zerolog.New(os.Stdout)

	ctx := logger.WithContext(context.Background())

	router := gin.Default()

	router.Use(gin.LoggerWithWriter(logger))

	apiKey := os.Getenv("GEMINI_API_KEY")

	gemini, err := googleai.New(ctx, googleai.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("error on gemini api connect %v", err)
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	var req entity.ScriptBreakdownRequest

	validate.RegisterStructValidation(entity.ScriptBreakdownRequestValidate, req)

	router.POST("/script/breakdown", presentationbreakdown.New(
		scriptbreakdown.New(
			validate,
			finaldraft.New(),
			scenebreakdown.New(
				ai.New(gemini, cache.New()),
			),
		),
	).ProcessFile)

	err = router.Run(":9000")
	if err != nil {
		logger.Fatal().Msg(err.Error())
	}
}
