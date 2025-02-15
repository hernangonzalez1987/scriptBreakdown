package queue

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/_interfaces"
)

type SQSListener struct {
	sqsClient    *sqs.Client
	queueURL     string
	eventHandler _interfaces.EventHandler
}

func NewSQSListener(sqsClient *sqs.Client, queueURL string, eventHandler _interfaces.EventHandler) *SQSListener {
	return &SQSListener{
		sqsClient:    sqsClient,
		queueURL:     queueURL,
		eventHandler: eventHandler,
	}
}

func (l *SQSListener) Listen(ctx context.Context) error {
	for {
		output, err := l.sqsClient.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
			QueueUrl:            &l.queueURL,
			MaxNumberOfMessages: 10,
			WaitTimeSeconds:     20,
		})
		if err != nil {
			log.Printf("Error receiving SQS message: %v", err)
			continue
		}

		for _, message := range output.Messages {
			var event events.S3Event
			err := json.Unmarshal([]byte(*message.Body), &event)
			if err != nil {
				log.Printf("Error unmarshalling SQS message: %v", err)
				continue
			}

			err = l.eventHandler.HandleEvent(ctx, event)
			if err != nil {
				log.Printf("Error processing script breakdown: %v", err)
				continue
			}

			_, err = l.sqsClient.DeleteMessage(ctx, &sqs.DeleteMessageInput{
				QueueUrl:      &l.queueURL,
				ReceiptHandle: message.ReceiptHandle,
			})
			if err != nil {
				log.Printf("Error deleting SQS message: %v", err)
			}
		}
	}
}
