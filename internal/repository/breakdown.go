package repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/_interfaces"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	"github.com/pkg/errors"
)

const (
	tableName               = "ScriptBreakdownResults"
	keyName                 = "script-breakdown-id"
	tableAlreadyExistsError = "Cannot create preexisting table"
)

type Repository struct {
	client _interfaces.DB
}

func New(client _interfaces.DB) *Repository {
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
	if strings.Contains(err.Error(), tableAlreadyExistsError) {
		return nil
	}
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (ref *Repository) Save(ctx context.Context, result entity.ScriptBreakdownResult) error {
	condition := "attribute_not_exists(BreakdownID)"

	result.LastUpdate = time.Now()

	if result.Version > 1 {
		condition = fmt.Sprintf("Version = %v", result.Version-1)
	}
	_, err := ref.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]types.AttributeValue{
			keyName: &types.AttributeValueMemberS{Value: result.BreakdownID},
		},
		ConditionExpression: aws.String(condition),
	})
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (ref *Repository) Get(ctx context.Context, id string) (*entity.ScriptBreakdownResult, error) {
	item, err := ref.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			keyName: &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	result := &entity.ScriptBreakdownResult{}
	err = attributevalue.UnmarshalMap(item.Item, result)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return result, nil
}
