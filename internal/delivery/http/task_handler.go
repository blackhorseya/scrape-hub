package http

import (
	"net/http"

	"github.com/blackhorseya/scrape-hub/internal/usecase/query"
	"github.com/blackhorseya/scrape-hub/pkg/contextx"
	"github.com/gin-gonic/gin"
)

// TaskHandler 處理任務相關的 HTTP 請求
type TaskHandler struct {
	taskQuery *query.TaskQuery
	group     *gin.RouterGroup
}

// NewTaskHandler 建立任務處理器實例
func NewTaskHandler(group *gin.RouterGroup, taskQuery *query.TaskQuery) *TaskHandler {
	handler := &TaskHandler{
		taskQuery: taskQuery,
		group:     group,
	}

	handler.RegisterRoutes()
	return handler
}

// RegisterRoutes 註冊路由
func (h *TaskHandler) RegisterRoutes() {
	v1 := h.group.Group("/v1")
	{
		v1.GET("/tasks", h.ListScheduledTasks)
	}
}

// ListScheduledTasks godoc
// @Summary 列出由定時任務觸發的 Lambda 函數
// @Description 取得所有使用 cron 表達式排程的 Lambda 函數清單
// @Tags tasks
// @Produce json
// @Security BearerAuth
// @Success 200 {object} Response{data=[]entity.Task}
// @Failure 500 {object} Response
// @Router /v1/tasks [get]
func (h *TaskHandler) ListScheduledTasks(c *gin.Context) {
	ctx := contextx.WithContext(c)

	tasks, err := h.taskQuery.ListScheduledTasks(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Data: tasks,
	})
}
