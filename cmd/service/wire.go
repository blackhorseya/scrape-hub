//go:build wireinject
// +build wireinject

package main

import (
	"github.com/blackhorseya/scrape-hub/configs"
	"github.com/blackhorseya/scrape-hub/internal/app/providers"
	httpdelivery "github.com/blackhorseya/scrape-hub/internal/delivery/http"
	"github.com/google/wire"
)

// InitializeApp 初始化應用程式
func InitializeApp(cfg *configs.Config) (httpdelivery.Server, error) {
	wire.Build(
		// AWS 客戶端
		providers.AWSSet,
		// 儲存庫
		providers.RepositorySet,
		// 用例
		providers.UseCaseSet,
		// HTTP 伺服器
		providers.ServerSet,
	)
	return nil, nil
}
