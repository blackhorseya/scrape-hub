package persistence

import (
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/blackhorseya/scrape-hub/internal/domain/entity"
	"github.com/blackhorseya/scrape-hub/internal/domain/repository"
	"github.com/blackhorseya/scrape-hub/pkg/contextx"
)

// TaskRepositoryImpl 實作 TaskRepository 界面
type TaskRepositoryImpl struct {
	ebClient *eventbridge.Client
	lwClient *lambda.Client
	cwClient *cloudwatch.Client
}

// NewTaskRepository 建立任務儲存庫實例
func NewTaskRepository(
	ebClient *eventbridge.Client,
	lwClient *lambda.Client,
	cwClient *cloudwatch.Client,
) repository.TaskRepository {
	return &TaskRepositoryImpl{
		ebClient: ebClient,
		lwClient: lwClient,
		cwClient: cwClient,
	}
}

// ListTasksBySchedule 實作列出所有由定時任務觸發的 Lambda 函數
func (r *TaskRepositoryImpl) ListTasksBySchedule(ctx contextx.Contextx) ([]*entity.Task, error) {
	// 取得所有 EventBridge 規則
	rules, err := r.ebClient.ListRules(ctx, &eventbridge.ListRulesInput{})
	if err != nil {
		return nil, fmt.Errorf("列出 EventBridge 規則失敗: %w", err)
	}

	tasks := make([]*entity.Task, 0)
	for _, rule := range rules.Rules {
		// 過濾出使用 cron 表達式的規則
		if !strings.HasPrefix(*rule.ScheduleExpression, "cron") {
			continue
		}

		// 取得規則的目標
		targets, err := r.ebClient.ListTargetsByRule(ctx, &eventbridge.ListTargetsByRuleInput{
			Rule: rule.Name,
		})
		if err != nil {
			ctx.Warn("取得規則目標失敗", "rule", *rule.Name, "error", err)
			continue
		}

		for _, target := range targets.Targets {
			// 確保目標是 Lambda 函數
			if !strings.Contains(*target.Arn, ":lambda:") {
				continue
			}

			// 從 ARN 中提取函數名稱
			parts := strings.Split(*target.Arn, ":")
			functionName := parts[len(parts)-1]

			// 取得函數的最後執行狀態
			status, lastTriggerTime := r.getLastExecutionStatus(ctx, functionName)

			task := &entity.Task{
				FunctionName:        functionName,
				CronExpression:      *rule.ScheduleExpression,
				LastTriggeredTime:   lastTriggerTime,
				LastExecutionStatus: status,
			}
			tasks = append(tasks, task)
		}
	}

	return tasks, nil
}

// getLastExecutionStatus 取得函數的最後執行狀態和時間
func (r *TaskRepositoryImpl) getLastExecutionStatus(ctx contextx.Contextx, functionName string) (entity.TaskStatus, time.Time) {
	// 設定時間範圍為過去 24 小時
	endTime := time.Now()
	startTime := endTime.Add(-24 * time.Hour)

	// 取得 CloudWatch 指標
	metrics, err := r.cwClient.GetMetricData(ctx, &cloudwatch.GetMetricDataInput{
		StartTime: &startTime,
		EndTime:   &endTime,
		MetricDataQueries: []types.MetricDataQuery{
			{
				Id: aws.String("errors"),
				MetricStat: &types.MetricStat{
					Metric: &types.Metric{
						Namespace:  aws.String("AWS/Lambda"),
						MetricName: aws.String("Errors"),
						Dimensions: []types.Dimension{
							{
								Name:  aws.String("FunctionName"),
								Value: aws.String(functionName),
							},
						},
					},
					Period: aws.Int32(60),
					Stat:   aws.String("Sum"),
				},
			},
			{
				Id: aws.String("invocations"),
				MetricStat: &types.MetricStat{
					Metric: &types.Metric{
						Namespace:  aws.String("AWS/Lambda"),
						MetricName: aws.String("Invocations"),
						Dimensions: []types.Dimension{
							{
								Name:  aws.String("FunctionName"),
								Value: aws.String(functionName),
							},
						},
					},
					Period: aws.Int32(60),
					Stat:   aws.String("Sum"),
				},
			},
		},
	})

	if err != nil {
		ctx.Warn("取得指標資料失敗", "function", functionName, "error", err)
		return entity.TaskStatusFailure, time.Time{}
	}

	// 如果有資料點，檢查最後一次執行是否有錯誤
	if len(metrics.MetricDataResults) > 1 {
		errors := metrics.MetricDataResults[0]
		invocations := metrics.MetricDataResults[1]

		if len(invocations.Timestamps) > 0 {
			lastIndex := len(invocations.Timestamps) - 1
			lastTime := invocations.Timestamps[lastIndex]

			// 檢查最後一次執行是否有錯誤
			if len(errors.Values) > lastIndex && errors.Values[lastIndex] > 0 {
				return entity.TaskStatusFailure, lastTime
			}
			return entity.TaskStatusSuccess, lastTime
		}
	}

	return entity.TaskStatusSuccess, time.Time{}
}
