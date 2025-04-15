package query

import (
	"github.com/blackhorseya/scrape-hub/internal/domain/entity"
	"github.com/blackhorseya/scrape-hub/internal/domain/repository"
	"github.com/blackhorseya/scrape-hub/pkg/contextx"
)

// TaskQuery 定義任務查詢用例
type TaskQuery struct {
	taskRepo repository.TaskRepository
}

// NewTaskQuery 建立任務查詢用例
func NewTaskQuery(taskRepo repository.TaskRepository) *TaskQuery {
	return &TaskQuery{taskRepo: taskRepo}
}

// ListScheduledTasks 列出所有由定時任務觸發的 Lambda 函數
func (q *TaskQuery) ListScheduledTasks(ctx contextx.Contextx) ([]*entity.Task, error) {
	tasks, err := q.taskRepo.ListTasksBySchedule(ctx)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
