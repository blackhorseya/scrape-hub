import { FC } from 'react';

interface TaskStatus {
  status: 'running' | 'success' | 'failed';
  count: number;
  label: string;
  color: string;
}

interface TaskStatusSummaryCardProps {
  status: TaskStatus;
}

// 模擬資料
export const mockTaskStatuses: TaskStatus[] = [
  { status: 'running', count: 5, label: '執行中', color: 'bg-blue-500' },
  { status: 'success', count: 42, label: '成功', color: 'bg-green-500' },
  { status: 'failed', count: 3, label: '失敗', color: 'bg-red-500' },
];

const TaskStatusSummaryCard: FC<TaskStatusSummaryCardProps> = ({ status }) => {
  return (
    <div className={`p-6 rounded-lg shadow-sm border border-gray-200 dark:border-gray-800`}>
      <div className="flex items-center">
        <div className={`w-4 h-4 rounded-full ${status.color} mr-3`} />
        <h3 className="text-lg font-semibold">{status.label}</h3>
      </div>
      <p className="mt-2 text-3xl font-bold">{status.count}</p>
    </div>
  );
};

export default TaskStatusSummaryCard;