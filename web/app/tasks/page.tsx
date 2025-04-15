import { FC } from 'react';
import TaskList from '@/components/tasks/TaskList';

const TasksPage: FC = () => {
  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <h1 className="text-2xl font-bold">任務管理</h1>
      </div>
      
      <TaskList />
    </div>
  );
};

export default TasksPage;