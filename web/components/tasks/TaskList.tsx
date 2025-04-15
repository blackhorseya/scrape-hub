import { FC } from 'react';

// 任務狀態型別
export interface Task {
  functionName: string;
  cronExpression: string;
  lastTriggeredTime: string;
  lastExecutionStatus: 'Success' | 'Failure';
}

// 模擬資料
const mockTasks: Task[] = [
  {
    functionName: '新聞爬蟲',
    cronExpression: '0 */4 * * *',
    lastTriggeredTime: '2025-04-15T08:00:00Z',
    lastExecutionStatus: 'Success',
  },
  {
    functionName: '天氣資料爬蟲',
    cronExpression: '0 0 * * *',
    lastTriggeredTime: '2025-04-15T00:00:00Z',
    lastExecutionStatus: 'Success',
  },
  {
    functionName: '股市資料爬蟲',
    cronExpression: '*/30 9-16 * * 1-5',
    lastTriggeredTime: '2025-04-15T06:30:00Z',
    lastExecutionStatus: 'Failure',
  },
  {
    functionName: '社群媒體爬蟲',
    cronExpression: '0 */2 * * *',
    lastTriggeredTime: '2025-04-15T08:00:00Z',
    lastExecutionStatus: 'Success',
  },
  {
    functionName: '商品價格爬蟲',
    cronExpression: '0 */6 * * *',
    lastTriggeredTime: '2025-04-15T06:00:00Z',
    lastExecutionStatus: 'Success',
  },
  {
    functionName: 'RSS 訂閱爬蟲',
    cronExpression: '*/15 * * * *',
    lastTriggeredTime: '2025-04-15T08:15:00Z',
    lastExecutionStatus: 'Success',
  }
];

// cron 表達式的中文說明
const getCronDescription = (expression: string): string => {
  if (expression === '0 */4 * * *') return '每 4 小時';
  if (expression === '0 0 * * *') return '每天凌晨';
  if (expression === '*/30 9-16 * * 1-5') return '工作日 9-16 點每 30 分鐘';
  if (expression === '0 */2 * * *') return '每 2 小時';
  if (expression === '0 */6 * * *') return '每 6 小時';
  if (expression === '*/15 * * * *') return '每 15 分鐘';
  return expression;
};

const TaskList: FC = () => {
  const formatDate = (dateString: string) => {
    const date = new Date(dateString);
    return new Intl.DateTimeFormat('zh-TW', {
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
      hour: '2-digit',
      minute: '2-digit',
      second: '2-digit',
      hour12: false
    }).format(date);
  };

  return (
    <div className="relative overflow-x-auto shadow-md sm:rounded-lg">
      <table className="w-full text-sm text-left text-gray-500 dark:text-gray-400">
        <thead className="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
          <tr>
            <th scope="col" className="px-6 py-3">任務名稱</th>
            <th scope="col" className="px-6 py-3">執行頻率</th>
            <th scope="col" className="px-6 py-3">最後執行時間</th>
            <th scope="col" className="px-6 py-3">狀態</th>
          </tr>
        </thead>
        <tbody>
          {mockTasks.map((task, index) => (
            <tr 
              key={task.functionName} 
              className={`${index % 2 === 0 ? 'bg-white' : 'bg-gray-50'} border-b dark:border-gray-700 dark:bg-gray-800 hover:bg-gray-100 dark:hover:bg-gray-700`}
            >
              <td className="px-6 py-4 font-medium text-gray-900 dark:text-white whitespace-nowrap">
                {task.functionName}
              </td>
              <td className="px-6 py-4" title={task.cronExpression}>
                {getCronDescription(task.cronExpression)}
              </td>
              <td className="px-6 py-4">
                {formatDate(task.lastTriggeredTime)}
              </td>
              <td className="px-6 py-4">
                <span className={`px-2.5 py-0.5 text-xs font-medium rounded-full ${
                  task.lastExecutionStatus === 'Success' 
                    ? 'bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-300'
                    : 'bg-red-100 text-red-800 dark:bg-red-900 dark:text-red-300'
                }`}>
                  {task.lastExecutionStatus === 'Success' ? '成功' : '失敗'}
                </span>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default TaskList;