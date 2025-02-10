package repository

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/_interfaces"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	"github.com/pkg/errors"
)

const tableName = "ScriptBreakdownResults"

type Repository struct {
	client _interfaces.DB
}

func New(client _interfaces.DB) *Repository {
	return &Repository{client: client}
}

func (ref *Repository) Save(ctx context.Context, result entity.ScriptBreakdownResult) error {
	condition := "attribute_not_exists(BreakdownID)"

	if result.Version > 1 {
		condition = fmt.Sprintf("Version = %v", result.Version-1)
	}
	_, err := ref.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]types.AttributeValue{
			"BreakdownID": &types.AttributeValueMemberS{Value: result.BreakdownID},
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
		Key: map[string]types.AttributeValue{
			"BreakdownID": &types.AttributeValueMemberS{Value: id},
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
