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
	"github.com/hernangonzalez1987/scriptBreakdown/internal/_interfaces"
	scriptbreakdownrequest "github.com/hernangonzalez1987/scriptBreakdown/internal/domain/useCase/scriptBreakdownRequest"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/useCase/scriptbreakdown"
	scenebreakdown "github.com/hernangonzalez1987/scriptBreakdown/internal/domain/useCase/scriptbreakdown/sceneBreakdown"
	finaldraft "github.com/hernangonzalez1987/scriptBreakdown/internal/integration/finalDraft"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/integration/llm"
	presentationbreakdown "github.com/hernangonzalez1987/scriptBreakdown/internal/presentation/breakdown"
	eventhandler "github.com/hernangonzalez1987/scriptBreakdown/internal/presentation/eventHandler"
	breakdownresultrepository "github.com/hernangonzalez1987/scriptBreakdown/internal/repository/breakdownResult"
	sceneanalysisrepository "github.com/hernangonzalez1987/scriptBreakdown/internal/repository/sceneAnalysis"
	"github.com/hernangonzalez1987/scriptBreakdown/tools/cache"
	"github.com/hernangonzalez1987/scriptBreakdown/tools/infrastructure/queue"
	"github.com/hernangonzalez1987/scriptBreakdown/tools/infrastructure/storage"
	"github.com/hernangonzalez1987/scriptBreakdown/tools/logger"
	uuidgenerator "github.com/hernangonzalez1987/scriptBreakdown/tools/uuidGenerator"
	"github.com/pkg/errors"
	"github.com/tmc/langchaingo/llms/googleai"
)

func main() {
	logger := logger.New()

	ctx := logger.AssociateWithContext(context.Background())

	router := gin.Default()

	router.Use(gin.LoggerWithWriter(logger.Logger()))

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

	breakdownResultRepository, sceneAnalysisRepository, err := getRepositories(ctx, awsConfig)
	if err != nil {
		log.Fatalf("error creating repositories %v", err)
	}

	sourceStorage, targetStorage := getStorages(awsConfig)

	queueClient := getSQSClient(awsConfig)

	router.POST("/script/breakdown", presentationbreakdown.New(
		scriptbreakdownrequest.New(sourceStorage, breakdownResultRepository),
	).ProcessFile)

	router.GET("/script/breakdown/:breakdownID", presentationbreakdown.New(
		scriptbreakdownrequest.New(sourceStorage, breakdownResultRepository),
	).GetResult)

	go func() {
		err := queue.NewSQSListener(queueClient, os.Getenv("BREAKDOWN_EVENTS_QUEUE"), eventhandler.NewS3EventHandler(
			scriptbreakdown.New(
				finaldraft.NewParser(),
				finaldraft.NewRender(),
				scenebreakdown.New(llm.New(gemini, cache.New[string](&ttl)), uuidgenerator.New(),
					sceneAnalysisRepository),
				sourceStorage,
				targetStorage,
				breakdownResultRepository,
			))).Listen(ctx)
		if err != nil {
			logger.Logger().Fatal().Msg(err.Error())
		}
	}()

	err = router.Run(":9000")
	if err != nil {
		logger.Logger().Fatal().Msg(err.Error())
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

func getRepositories(ctx context.Context, awsConfig aws.Config) (
	_interfaces.BreakdownRepository,
	_interfaces.SceneAnalysisRepository,
	error,
) {
	dbClient := getDynamoClient(awsConfig)

	breakdownResultRepository := breakdownresultrepository.New(dbClient)
	sceneAnalysisRepository := sceneanalysisrepository.New(dbClient)

	err := breakdownResultRepository.Init(ctx)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	err = sceneAnalysisRepository.Init(ctx)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	return breakdownResultRepository, sceneAnalysisRepository, nil
}

func getStorages(awsConfig aws.Config) (*storage.S3Storage, *storage.S3Storage) {
	storageClient := getS3Client(awsConfig)

	return storage.NewS3Storage(storageClient, os.Getenv("SCRIPTS_BUCKET")),
		storage.NewS3Storage(storageClient, os.Getenv("BREAKDOWNS_BUCKET"))
}
