import { FC } from 'react';
import TaskStatusSummaryCard, { mockTaskStatuses } from '@/components/dashboard/TaskStatusSummaryCard';
import RecentExecutionsTable from '@/components/dashboard/RecentExecutionsTable';
import StatusTrendChart from '@/components/dashboard/StatusTrendChart';

const Home: FC = () => {
  return (
    <div className="space-y-6">
      <h1 className="text-2xl font-bold">儀表板</h1>
      
      {/* 狀態統計卡片 */}
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        {mockTaskStatuses.map((status) => (
          <TaskStatusSummaryCard key={status.status} status={status} />
        ))}
      </div>

      {/* 趨勢圖表 */}
      <div className="mt-6">
        <StatusTrendChart />
      </div>

      {/* 最近執行記錄 */}
      <div className="mt-6">
        <h2 className="text-xl font-semibold mb-4">最近執行記錄</h2>
        <RecentExecutionsTable />
      </div>
    </div>
  );
};

export default Home;
