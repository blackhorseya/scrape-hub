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
}

// NewTaskHandler 建立任務處理器實例
func NewTaskHandler(taskQuery *query.TaskQuery) *TaskHandler {
	return &TaskHandler{
		taskQuery: taskQuery,
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
		ctx.Error("列出排程任務失敗", "error", err)
		c.JSON(http.StatusInternalServerError, NewErrorResponse(
			http.StatusInternalServerError,
			"取得任務列表失敗",
		))
		return
	}

	c.JSON(http.StatusOK, NewSuccessResponse(tasks))
}
