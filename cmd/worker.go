/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/useCase/scriptbreakdown"
	scenebreakdown "github.com/hernangonzalez1987/scriptBreakdown/internal/domain/useCase/scriptbreakdown/sceneBreakdown"
	finaldraft "github.com/hernangonzalez1987/scriptBreakdown/internal/integration/finalDraft"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/integration/llm"
	eventhandler "github.com/hernangonzalez1987/scriptBreakdown/internal/presentation/eventHandler"
	"github.com/hernangonzalez1987/scriptBreakdown/tools/cache"
	"github.com/hernangonzalez1987/scriptBreakdown/tools/infrastructure/queue"
	"github.com/hernangonzalez1987/scriptBreakdown/tools/logger"
	uuidgenerator "github.com/hernangonzalez1987/scriptBreakdown/tools/uuidGenerator"
	"github.com/spf13/cobra"
	"github.com/tmc/langchaingo/llms/googleai"
)

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Start Script Breakdown Worker",
	Long: `This worker process async the script breakdown requests
	created on the API`,
	Run: func(_ *cobra.Command, _ []string) {
		apiKey := os.Getenv("GEMINI_API_KEY")

		logger := logger.New()

		ctx := logger.AssociateWithContext(context.Background())

		gemini, err := googleai.New(ctx, googleai.WithAPIKey(apiKey), googleai.WithDefaultModel("gemini-1.5-flash"))
		if err != nil {
			log.Fatalf("error on gemini api connect %v", err)
		}

		awsConfig, err := config.LoadDefaultConfig(ctx, config.WithSharedConfigProfile("local"))
		if err != nil {
			log.Fatalf("error creating aws config %v", err)
		}

		queueClient := getSQSClient(awsConfig)

		scriptsStorage, breakdownsStorage := getStorages(awsConfig)

		ttl := time.Hour

		breakdownResultRepository, sceneAnalysisRepository, err := getRepositories(ctx, awsConfig)
		if err != nil {
			log.Fatalf("error getting repositories %v", err)
		}

		err = queue.NewSQSListener(queueClient, os.Getenv("BREAKDOWN_EVENTS_QUEUE"), eventhandler.NewS3EventHandler(
			scriptbreakdown.New(
				finaldraft.NewParser(),
				finaldraft.NewRender(),
				scenebreakdown.New(llm.New(gemini, cache.New[string](&ttl)), uuidgenerator.New(),
					sceneAnalysisRepository),
				scriptsStorage,
				breakdownsStorage,
				breakdownResultRepository,
			))).Listen(ctx)
		if err != nil {
			logger.Logger().Fatal().Msg(err.Error())
		}
	},
}

func init() {
	startCmd.AddCommand(workerCmd)
}
