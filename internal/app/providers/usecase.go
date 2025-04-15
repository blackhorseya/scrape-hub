package providers

import (
	"github.com/blackhorseya/scrape-hub/internal/domain/repository"
	"github.com/blackhorseya/scrape-hub/internal/usecase/query"
	"github.com/google/wire"
)

// ProvideTaskQuery 提供任務查詢用例實例
func ProvideTaskQuery(taskRepo repository.TaskRepository) *query.TaskQuery {
	return query.NewTaskQuery(taskRepo)
}

// UseCaseSet 提供用例相關的依賴注入集合
var UseCaseSet = wire.NewSet(ProvideTaskQuery)
