package repository

import (
	"github.com/blackhorseya/scrape-hub/internal/domain/entity"
	"github.com/blackhorseya/scrape-hub/pkg/contextx"
)

// TaskRepository 定義任務儲存庫的操作介面
type TaskRepository interface {
	// ListTasksBySchedule 列出所有由定時任務觸發的 Lambda 函數
	ListTasksBySchedule(ctx contextx.Contextx) ([]*entity.Task, error)
}
