package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	lambdasvc "github.com/aws/aws-sdk-go-v2/service/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/blackhorseya/scrape-hub/configs"
	httpdelivery "github.com/blackhorseya/scrape-hub/internal/delivery/http"
	"github.com/blackhorseya/scrape-hub/internal/domain/repository"
	"github.com/blackhorseya/scrape-hub/internal/infra/persistence"
	"github.com/blackhorseya/scrape-hub/internal/usecase/query"
)

func initAWSClients() (*eventbridge.Client, *lambdasvc.Client, *cloudwatch.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, nil, nil, fmt.Errorf("載入 AWS 設定失敗: %w", err)
	}

	ebClient := eventbridge.NewFromConfig(cfg)
	lwClient := lambdasvc.NewFromConfig(cfg)
	cwClient := cloudwatch.NewFromConfig(cfg)

	return ebClient, lwClient, cwClient, nil
}

func initRepositories(ebClient *eventbridge.Client, lwClient *lambdasvc.Client, cwClient *cloudwatch.Client) (repository.TaskRepository, error) {
	return persistence.NewTaskRepository(ebClient, lwClient, cwClient), nil
}

func initUseCases(taskRepo repository.TaskRepository) *query.TaskQuery {
	return query.NewTaskQuery(taskRepo)
}

// initServer 初始化 HTTP 伺服器並註冊所有路由
func initServer(cfg *configs.Config, taskQuery *query.TaskQuery) (httpdelivery.Server, error) {
	server, err := httpdelivery.NewServer(cfg)
	if err != nil {
		return nil, fmt.Errorf("初始化伺服器失敗: %w", err)
	}

	// 註冊任務處理器
	engine := server.Engine()
	apiGroup := engine.Group("/api")
	httpdelivery.NewTaskHandler(apiGroup, taskQuery)

	return server, nil
}

func main() {
	cfg, err := configs.LoadFromEnv()
	if err != nil {
		log.Fatalf("載入設定失敗: %v", err)
	}

	// 初始化各層元件
	ebClient, lwClient, cwClient, err := initAWSClients()
	if err != nil {
		log.Fatalf("初始化 AWS 客戶端失敗: %v", err)
	}

	taskRepo, err := initRepositories(ebClient, lwClient, cwClient)
	if err != nil {
		log.Fatalf("初始化儲存庫失敗: %v", err)
	}

	taskQuery := initUseCases(taskRepo)

	server, err := initServer(cfg, taskQuery)
	if err != nil {
		log.Fatalf("初始化伺服器失敗: %v", err)
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
