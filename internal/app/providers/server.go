package providers

import (
	"github.com/blackhorseya/scrape-hub/configs"
	httpdelivery "github.com/blackhorseya/scrape-hub/internal/delivery/http"
	"github.com/blackhorseya/scrape-hub/internal/usecase/query"
	"github.com/google/wire"
)

// ProvideHTTPServer 提供 HTTP 伺服器實例
func ProvideHTTPServer(cfg *configs.Config, taskQuery *query.TaskQuery) (httpdelivery.Server, error) {
	server, err := httpdelivery.NewServer(cfg)
	if err != nil {
		return nil, err
	}

	// 註冊任務處理器
	apiGroup := server.Engine().Group("/api")
	httpdelivery.NewTaskHandler(apiGroup, taskQuery)

	return server, nil
}

// ServerSet 提供伺服器相關的依賴注入集合
var ServerSet = wire.NewSet(ProvideHTTPServer)
