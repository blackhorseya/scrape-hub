import { FC } from 'react';
import Link from 'next/link';

interface Execution {
  id: string;
  taskName: string;
  status: 'running' | 'success' | 'failed';
  startTime: string;
  duration: string;
}

// 模擬資料
export const mockExecutions: Execution[] = [
  {
    id: '1',
    taskName: '商品價格爬蟲',
    status: 'success',
    startTime: '2025-04-14 09:00:00',
    duration: '2分鐘'
  },
  {
    id: '2',
    taskName: '新聞標題爬蟲',
    status: 'running',
    startTime: '2025-04-14 09:15:00',
    duration: '進行中'
  },
  {
    id: '3',
    taskName: '股票資訊爬蟲',
    status: 'failed',
    startTime: '2025-04-14 09:10:00',
    duration: '1分鐘'
  },
  {
    id: '4',
    taskName: '天氣資訊爬蟲',
    status: 'success',
    startTime: '2025-04-14 09:05:00',
    duration: '3分鐘'
  },
  {
    id: '5',
    taskName: '論壇文章爬蟲',
    status: 'success',
    startTime: '2025-04-14 09:00:00',
    duration: '5分鐘'
  }
];

const statusStyles = {
  running: 'bg-blue-100 text-blue-800 dark:bg-blue-900 dark:text-blue-300',
  success: 'bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-300',
  failed: 'bg-red-100 text-red-800 dark:bg-red-900 dark:text-red-300'
};

const statusLabels = {
  running: '執行中',
  success: '成功',
  failed: '失敗'
};

const RecentExecutionsTable: FC = () => {
  return (
    <div className="overflow-x-auto">
      <table className="min-w-full divide-y divide-gray-200 dark:divide-gray-800">
        <thead className="bg-gray-50 dark:bg-gray-900">
          <tr>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">
              任務名稱
            </th>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">
              狀態
            </th>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">
              開始時間
            </th>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">
              執行時間
            </th>
          </tr>
        </thead>
        <tbody className="bg-white dark:bg-gray-900 divide-y divide-gray-200 dark:divide-gray-800">
          {mockExecutions.map((execution) => (
            <tr key={execution.id}>
              <td className="px-6 py-4 whitespace-nowrap text-sm">
                <Link 
                  href={`/tasks/${execution.id}`}
                  className="text-blue-600 dark:text-blue-400 hover:underline"
                >
                  {execution.taskName}
                </Link>
              </td>
              <td className="px-6 py-4 whitespace-nowrap text-sm">
                <span className={`px-2 inline-flex text-xs leading-5 font-semibold rounded-full ${statusStyles[execution.status]}`}>
                  {statusLabels[execution.status]}
                </span>
              </td>
              <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500 dark:text-gray-400">
                {execution.startTime}
              </td>
              <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500 dark:text-gray-400">
                {execution.duration}
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default RecentExecutionsTable;