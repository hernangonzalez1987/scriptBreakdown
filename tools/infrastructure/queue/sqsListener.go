package queue

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/hernangonzalez1987/scriptBreakdown/tools/logger"
	log "github.com/rs/zerolog"
)

const (
	maxNumberOfMessages = 10
	waitTimeSeconds     = 20
)

type SQSListener struct {
	sqsClient    *sqs.Client
	queueURL     string
	eventHandler EventHandler
}

func NewSQSListener(sqsClient *sqs.Client, queueURL string, eventHandler EventHandler) *SQSListener {
	return &SQSListener{
		sqsClient:    sqsClient,
		queueURL:     queueURL,
		eventHandler: eventHandler,
	}
}

func (l *SQSListener) Listen(ctx context.Context) error {
	for {
		log.Ctx(ctx).Info().Msg("listening for events")

		output, err := l.sqsClient.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
			QueueUrl:            &l.queueURL,
			MaxNumberOfMessages: maxNumberOfMessages,
			WaitTimeSeconds:     waitTimeSeconds,
		})
		if err != nil {
			log.Ctx(ctx).Error().Err(err).Msg("error receiving from queue")

			continue
		}

		for _, message := range output.Messages {
			var event events.S3Event

			err := json.Unmarshal([]byte(*message.Body), &event)
			if err != nil {
				log.Ctx(ctx).Warn().Err(err).Msg("error unmarshaling event")

				continue
			}

			newCtx := logger.AddCorrelationID(ctx)

			err = l.eventHandler.HandleEvent(newCtx, event)
			if err != nil {
				log.Ctx(newCtx).Error().Err(err).Msg("error on event handler")

				continue
			}

			_, err = l.sqsClient.DeleteMessage(newCtx, &sqs.DeleteMessageInput{
				QueueUrl:      &l.queueURL,
				ReceiptHandle: message.ReceiptHandle,
			})
			if err != nil {
				log.Ctx(newCtx).Error().Err(err).Msg("error on delete from queue")
			}
		}
	}
}
