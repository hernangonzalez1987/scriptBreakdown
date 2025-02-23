package sceneanalysisrepository

import (
	"context"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	"github.com/pkg/errors"
)

const (
	tableName               = "SceneAnalysis"
	keyName                 = "scene_id"
	tableAlreadyExistsError = "Cannot create preexisting table"
)

type Repository struct {
	client *dynamodb.Client
}

func New(client *dynamodb.Client) *Repository {
	return &Repository{client: client}
}

func (ref *Repository) Init(ctx context.Context) error {
	_, err := ref.client.CreateTable(ctx, &dynamodb.CreateTableInput{
		TableName: aws.String(tableName),
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String(keyName),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String(keyName),
				KeyType:       types.KeyTypeHash,
			},
		},
		BillingMode: types.BillingModePayPerRequest,
	})
	if err != nil {
		if strings.Contains(err.Error(), tableAlreadyExistsError) {
			return nil
		}
		return errors.WithStack(err)
	}

	return nil
}

func (ref *Repository) Save(ctx context.Context, analysis entity.SceneAnalysis) error {
	analysis.UpdatedAt = time.Now()

	attrs, err := attributevalue.MarshalMap(analysis)
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = ref.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      attrs,
	})
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (ref *Repository) Get(ctx context.Context, id string) (*entity.SceneAnalysis, error) {
	item, err := ref.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			keyName: &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if item.Item == nil {
		return nil, nil
	}

	result := &entity.SceneAnalysis{}
	err = attributevalue.UnmarshalMap(item.Item, result)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return result, nil
}
