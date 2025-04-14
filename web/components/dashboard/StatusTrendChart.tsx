import { FC } from 'react';

interface DayStats {
  date: string;
  success: number;
  failed: number;
}

// 模擬過去 7 天的資料
export const mockTrendData: DayStats[] = [
  { date: '04/08', success: 38, failed: 2 },
  { date: '04/09', success: 45, failed: 1 },
  { date: '04/10', success: 41, failed: 4 },
  { date: '04/11', success: 39, failed: 3 },
  { date: '04/12', success: 42, failed: 2 },
  { date: '04/13', success: 40, failed: 5 },
  { date: '04/14', success: 42, failed: 3 },
];

const StatusTrendChart: FC = () => {
  const maxValue = Math.max(
    ...mockTrendData.map(day => Math.max(day.success, day.failed))
  );

  return (
    <div className="p-6 border border-gray-200 dark:border-gray-800 rounded-lg">
      <h3 className="text-lg font-semibold mb-4">執行狀態趨勢</h3>
      <div className="flex items-end space-x-2 h-40">
        {mockTrendData.map((day) => (
          <div key={day.date} className="flex-1 flex flex-col items-center">
            <div className="w-full flex flex-col items-center space-y-1">
              <div 
                className="w-full bg-red-500 rounded-t"
                style={{ height: `${(day.failed / maxValue) * 100}%` }}
                title={`失敗: ${day.failed}`}
              />
              <div
                className="w-full bg-green-500 rounded-b"
                style={{ height: `${(day.success / maxValue) * 100}%` }}
                title={`成功: ${day.success}`}
              />
            </div>
            <span className="mt-2 text-xs text-gray-500">{day.date}</span>
          </div>
        ))}
      </div>
      <div className="flex justify-center mt-4 space-x-4">
        <div className="flex items-center">
          <div className="w-3 h-3 bg-green-500 rounded mr-2" />
          <span className="text-sm text-gray-600 dark:text-gray-400">成功</span>
        </div>
        <div className="flex items-center">
          <div className="w-3 h-3 bg-red-500 rounded mr-2" />
          <span className="text-sm text-gray-600 dark:text-gray-400">失敗</span>
        </div>
      </div>
    </div>
  );
};

export default StatusTrendChart;