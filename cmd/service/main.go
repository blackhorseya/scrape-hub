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
	"github.com/blackhorseya/scrape-hub/internal/delivery/middleware"
	"github.com/gin-gonic/gin"
)

var auth0Middleware *middleware.Auth0Middleware

// initRouter 初始化並回傳 Gin 路由器
func initRouter(cfg *configs.Config) (*gin.Engine, error) {
	r := gin.Default()

	// 建立 Auth0 中介層
	var err error
	auth0Middleware, err = middleware.NewAuth0Middleware(&cfg.Auth0)
	if err != nil {
		return nil, fmt.Errorf("初始化 Auth0 中介層失敗: %w", err)
	}

	// 公開路由
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// 受保護的路由群組
	protected := r.Group("/api")
	protected.Use(auth0Middleware.EnsureValidToken())
	{
		protected.GET("/protected", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "這是一個受保護的端點",
			})
		})
	}

	return r, nil
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

	r, err := initRouter(cfg)
	if err != nil {
		log.Fatalf("初始化路由器失敗: %v", err)
	}

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
