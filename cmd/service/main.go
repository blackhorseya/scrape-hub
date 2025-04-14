package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/blackhorseya/scrape-hub/configs"
	"github.com/gin-gonic/gin"
)

// initRouter 初始化並回傳 Gin 路由器
func initRouter() *gin.Engine {
	r := gin.Default()

	// 設定路由
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	return r
}

// initLambdaHandler 初始化 Lambda 處理器
func initLambdaHandler(r *gin.Engine) *ginadapter.GinLambdaV2 {
	return ginadapter.NewV2(r)
}

// 載入設定並回傳
func loadConfig() *configs.Config {
	cfg, err := configs.LoadFromEnv()
	if err != nil {
		log.Fatalf("載入設定失敗: %v", err)
	}
	return cfg
}

func main() {
	cfg := loadConfig()
	r := initRouter()

	// 透過環境變數判斷執行環境
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != "" {
		// AWS Lambda 環境
		h := initLambdaHandler(r)
		lambda.Start(func(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
			return h.ProxyWithContext(ctx, req)
		})
	} else {
		// 本地開發環境
		addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
		log.Printf("本地伺服器啟動於 %s", addr)
		if err := r.Run(addr); err != nil {
			log.Fatalf("伺服器啟動失敗: %v", err)
		}
	}
}
