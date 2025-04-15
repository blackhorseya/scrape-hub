package providers

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	lambdasvc "github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/google/wire"
)

// AWSClients 包裝 AWS 客戶端
type AWSClients struct {
	EventBridge *eventbridge.Client
	Lambda      *lambdasvc.Client
	CloudWatch  *cloudwatch.Client
}

// ProvideAWSClients 提供 AWS 客戶端
func ProvideAWSClients() (*AWSClients, error) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, fmt.Errorf("載入 AWS 設定失敗: %w", err)
	}

	ebClient := eventbridge.NewFromConfig(cfg)
	lwClient := lambdasvc.NewFromConfig(cfg)
	cwClient := cloudwatch.NewFromConfig(cfg)

	return &AWSClients{
		EventBridge: ebClient,
		Lambda:      lwClient,
		CloudWatch:  cwClient,
	}, nil
}

// AWSSet 提供 AWS 相關的依賴注入集合
var AWSSet = wire.NewSet(ProvideAWSClients)
