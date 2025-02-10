package repository

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/_mocks"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/_interfaces"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
)

func TestRepository_Save(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	mockDb := _mocks.NewMockDB(t)

	breakdownNew := entity.ScriptBreakdownResult{
		BreakdownID: "someID",
		Version:     1,
	}

	breakdownUpdated := entity.ScriptBreakdownResult{
		BreakdownID: "someOtherID",
		Version:     2,
	}

	mockDb.EXPECT().PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]types.AttributeValue{
			"BreakdownID": &types.AttributeValueMemberS{Value: "someID"},
		},
		ConditionExpression: aws.String("attribute_not_exists(BreakdownID)"),
	}).Return(nil, nil)

	mockDb.EXPECT().PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]types.AttributeValue{
			"BreakdownID": &types.AttributeValueMemberS{Value: "someOtherID"},
		},
		ConditionExpression: aws.String("Version = 1"),
	}).Return(nil, nil)

	type fields struct {
		client _interfaces.DB
	}
	type args struct {
		ctx    context.Context
		result entity.ScriptBreakdownResult
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "should save new item",
			fields:  fields{client: mockDb},
			args:    args{ctx: ctx, result: breakdownNew},
			wantErr: false,
		},
		{
			name:    "should udpate item",
			fields:  fields{client: mockDb},
			args:    args{ctx: ctx, result: breakdownUpdated},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ref := &Repository{
				client: tt.fields.client,
			}
			if err := ref.Save(tt.args.ctx, tt.args.result); (err != nil) != tt.wantErr {
				t.Errorf("Repository.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
