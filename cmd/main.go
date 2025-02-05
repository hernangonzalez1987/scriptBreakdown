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
	"github.com/hernangonzalez1987/scriptBreakdown/internal/integration/ai/cache"
	finaldraft "github.com/hernangonzalez1987/scriptBreakdown/internal/integration/finalDraft"
	presentationbreakdown "github.com/hernangonzalez1987/scriptBreakdown/internal/presentation"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/tmc/langchaingo/llms/googleai"
)

func main() {

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

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

	validate.RegisterStructValidation(entity.ScriptBreakdownRequestValidate, entity.ScriptBreakdownRequest{})

	router.POST("/script/breakdown", presentationbreakdown.New(
		scriptbreakdown.New(
			validate,
			finaldraft.New(),
			scenebreakdown.New(
				ai.New(gemini, cache.New()),
			),
		),
	).ProcessFile)

	router.Run(":9000")
}
