/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/gin-gonic/gin"
	scriptbreakdownrequest "github.com/hernangonzalez1987/scriptBreakdown/internal/domain/useCase/scriptBreakdownRequest"
	presentationbreakdown "github.com/hernangonzalez1987/scriptBreakdown/internal/presentation/breakdown"
	"github.com/hernangonzalez1987/scriptBreakdown/tools/logger"
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Start Script Breakdown API",
	Long: `Start the Script Breakdown API.
	This API is used for creating Breakdown requests
	and getting the result.`,
	Run: func(_ *cobra.Command, _ []string) {
		logger := logger.New()

		ctx := logger.AssociateWithContext(context.Background())

		router := gin.Default()

		router.Use(gin.LoggerWithWriter(logger.Logger()))

		awsConfig, err := config.LoadDefaultConfig(ctx, config.WithSharedConfigProfile("local"))
		if err != nil {
			log.Fatalf("error creating aws config %v", err)
		}

		scriptsStorage, breakdownsStorage := getStorages(awsConfig)

		breakdownResultRepository, _, err := getRepositories(ctx, awsConfig)
		if err != nil {
			log.Fatalf("error creating aws config %v", err)
		}

		router.POST("/script/breakdown", presentationbreakdown.New(
			scriptbreakdownrequest.New(scriptsStorage, breakdownsStorage, breakdownResultRepository),
		).ProcessFile)

		router.GET("/script/breakdown/:breakdownID", presentationbreakdown.New(
			scriptbreakdownrequest.New(scriptsStorage, breakdownsStorage, breakdownResultRepository),
		).GetResult)

		err = router.Run(":9000")
		if err != nil {
			logger.Logger().Fatal().Msg(err.Error())
		}
	},
}

func init() {
	startCmd.AddCommand(apiCmd)
}
