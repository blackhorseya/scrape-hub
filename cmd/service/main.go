package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/blackhorseya/scrape-hub/configs"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambdaV2

func init() {
	// 載入設定
	_, err := configs.LoadFromEnv()
	if err != nil {
		log.Fatalf("載入設定失敗: %v", err)
	}

	// 初始化 Gin engine
	r := gin.Default()

	// 設定路由
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// 建立 Lambda 適配器
	ginLambda = ginadapter.NewV2(r)
}

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
