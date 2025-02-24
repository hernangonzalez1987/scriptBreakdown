package cmd

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	breakdownresultrepository "github.com/hernangonzalez1987/scriptBreakdown/internal/repository/breakdownResult"
	sceneanalysisrepository "github.com/hernangonzalez1987/scriptBreakdown/internal/repository/sceneAnalysis"
	"github.com/hernangonzalez1987/scriptBreakdown/tools/infrastructure/storage"
	"github.com/pkg/errors"
)

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
	*breakdownresultrepository.Repository,
	*sceneanalysisrepository.Repository,
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
