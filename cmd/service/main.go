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
)

func main() {
	cfg, err := configs.LoadFromEnv()
	if err != nil {
		log.Fatalf("載入設定失敗: %v", err)
	}

	server, err := InitializeApp(cfg)
	if err != nil {
		log.Fatalf("初始化應用程式失敗: %v", err)
	}

	// 透過環境變數判斷執行環境
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != "" {
		// AWS Lambda 環境
		h := ginadapter.NewV2(server.Engine())
		lambda.Start(func(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
			return h.ProxyWithContext(ctx, req)
		})
	} else {
		// 本地開發環境
		addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
		log.Printf("本地伺服器啟動於 %s", addr)
		if err := server.Engine().Run(addr); err != nil {
			log.Fatalf("伺服器啟動失敗: %v", err)
		}
	}
}
