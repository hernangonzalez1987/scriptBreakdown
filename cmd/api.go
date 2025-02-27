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

	_ "github.com/hernangonzalez1987/scriptBreakdown/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
		).BreakdownScript)

		router.GET("/script/breakdown/:breakdownID", presentationbreakdown.New(
			scriptbreakdownrequest.New(scriptsStorage, breakdownsStorage, breakdownResultRepository),
		).GetResult)

		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		err = router.Run(":9000")
		if err != nil {
			logger.Logger().Fatal().Msg(err.Error())
		}
	},
}

// @title           Script Breakdown API
// @version         1.0
// @description     This API allows to create and obtain results from script breakdowns

// @host      localhost:9000
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func init() {
	startCmd.AddCommand(apiCmd)
}
