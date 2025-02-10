package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/useCase/scriptbreakdown"
	scenebreakdown "github.com/hernangonzalez1987/scriptBreakdown/internal/domain/useCase/scriptbreakdown/sceneBreakdown"
	finaldraft "github.com/hernangonzalez1987/scriptBreakdown/internal/integration/finalDraft"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/integration/llm"
	presentationbreakdown "github.com/hernangonzalez1987/scriptBreakdown/internal/integration/llm/presentation"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/tools/cache"
	uuidgenerator "github.com/hernangonzalez1987/scriptBreakdown/internal/tools/uuidGenerator"
	"github.com/rs/zerolog"
	"github.com/tmc/langchaingo/llms/googleai"
)

func main() {
	logger := zerolog.New(os.Stdout).Level(zerolog.InfoLevel)

	ctx := logger.WithContext(context.Background())

	router := gin.Default()

	router.Use(gin.LoggerWithWriter(logger))

	apiKey := os.Getenv("GEMINI_API_KEY")

	gemini, err := googleai.New(ctx, googleai.WithAPIKey(apiKey), googleai.WithDefaultModel("gemini-1.5-flash"))
	if err != nil {
		log.Fatalf("error on gemini api connect %v", err)
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	var req entity.ScriptBreakdownRequest

	validate.RegisterStructValidation(entity.ScriptBreakdownRequestValidate, req)

	ttl := time.Hour

	router.POST("/script/breakdown", presentationbreakdown.New(
		scriptbreakdown.New(
			validate,
			finaldraft.NewParser(),
			finaldraft.NewRender(),
			scenebreakdown.New(
				llm.New(gemini, cache.New[string](&ttl)),
				uuidgenerator.New(),
			),
		),
	).ProcessFile)

	err = router.Run(":9000")
	if err != nil {
		logger.Fatal().Msg(err.Error())
	}
}
