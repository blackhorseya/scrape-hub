package entity

import "time"

// TaskStatus 定義任務執行狀態
type TaskStatus string

const (
	// TaskStatusSuccess 表示任務執行成功
	TaskStatusSuccess TaskStatus = "Success"
	// TaskStatusFailure 表示任務執行失敗
	TaskStatusFailure TaskStatus = "Failure"
)

// Task 代表一個定時任務
type Task struct {
	// FunctionName 是 Lambda 函數名稱
	FunctionName string `json:"functionName"`
	// CronExpression 是定時表達式
	CronExpression string `json:"cronExpression"`
	// LastTriggeredTime 是最後執行時間
	LastTriggeredTime time.Time `json:"lastTriggeredTime"`
	// LastExecutionStatus 是最後執行狀態
	LastExecutionStatus TaskStatus `json:"lastExecutionStatus"`
}
