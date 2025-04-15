package providers

import (
	"github.com/blackhorseya/scrape-hub/internal/domain/repository"
	"github.com/blackhorseya/scrape-hub/internal/infra/persistence"
	"github.com/google/wire"
)

// ProvideTaskRepository 提供任務儲存庫實例
func ProvideTaskRepository(clients *AWSClients) repository.TaskRepository {
	return persistence.NewTaskRepository(clients.EventBridge, clients.Lambda, clients.CloudWatch)
}

// RepositorySet 提供儲存庫相關的依賴注入集合
var RepositorySet = wire.NewSet(ProvideTaskRepository)
