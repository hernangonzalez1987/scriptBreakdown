package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/gin-gonic/gin"
	scriptbreakdownrequest "github.com/hernangonzalez1987/scriptBreakdown/internal/domain/useCase/scriptBreakdownRequest"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/useCase/scriptbreakdown"
	scenebreakdown "github.com/hernangonzalez1987/scriptBreakdown/internal/domain/useCase/scriptbreakdown/sceneBreakdown"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/infrastructure/queue"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/infrastructure/storage"
	finaldraft "github.com/hernangonzalez1987/scriptBreakdown/internal/integration/finalDraft"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/integration/llm"
	presentationbreakdown "github.com/hernangonzalez1987/scriptBreakdown/internal/presentation/breakdown"
	eventhandler "github.com/hernangonzalez1987/scriptBreakdown/internal/presentation/eventHandler"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/repository"
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

	ttl := time.Hour

	awsConfig, err := config.LoadDefaultConfig(ctx, config.WithSharedConfigProfile("local"))
	if err != nil {
		log.Fatalf("error creating aws config %v", err)
	}

	dbClient := getDynamoClient(awsConfig)
	queueClient := getSQSClient(awsConfig)
	storageClient := getS3Client(awsConfig)

	repository := repository.New(dbClient)

	err = repository.Init(ctx)
	if err != nil {
		log.Fatalf("error initializing repository %v", err)
	}

	sourceStorage := storage.NewS3Storage(storageClient, os.Getenv("SCRIPTS_BUCKET"))
	targetStorage := storage.NewS3Storage(storageClient, os.Getenv("BREAKDOWNS_BUCKET"))

	router.POST("/script/breakdown", presentationbreakdown.New(
		scriptbreakdownrequest.New(sourceStorage),
	).ProcessFile)

	go func() {
		err := queue.NewSQSListener(queueClient, os.Getenv("BREAKDOWN_EVENTS_QUEUE"), eventhandler.NewS3EventHandler(
			scriptbreakdown.New(
				finaldraft.NewParser(),
				finaldraft.NewRender(),
				scenebreakdown.New(llm.New(gemini, cache.New[string](&ttl)), uuidgenerator.New()),
				sourceStorage,
				targetStorage,
				repository,
			))).Listen(ctx)
		if err != nil {
			logger.Fatal().Msg(err.Error())
		}
	}()

	err = router.Run(":9000")
	if err != nil {
		logger.Fatal().Msg(err.Error())
	}
}

func getDynamoClient(cfg aws.Config) *dynamodb.Client {
	return dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
		o.BaseEndpoint = aws.String("http://localhost:8000")
	})
}

func getSQSClient(cfg aws.Config) *sqs.Client {
	return sqs.NewFromConfig(cfg)
}

func getS3Client(cfg aws.Config) *s3.Client {
	return s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})
}
