"use client";

import React from "react";
import Link from "next/link";
import { usePathname } from "next/navigation";
import { 
  FolderIcon, 
  HomeIcon, 
  ClockIcon,
  SettingsIcon,
  ChevronLeftIcon,
  ChevronRightIcon
} from "lucide-react";

interface SidebarProps {
  isExpanded: boolean;
  onToggle: () => void;
}

const navigation = [
  { name: "總覽", href: "/", icon: HomeIcon },
  { name: "爬蟲任務", href: "/tasks", icon: FolderIcon },
  { name: "執行記錄", href: "/history", icon: ClockIcon },
  { name: "系統設定", href: "/settings", icon: SettingsIcon },
];

export default function Sidebar({ isExpanded, onToggle }: SidebarProps) {
  const pathname = usePathname();
  
  return (
    <div className={`${isExpanded ? "w-64" : "w-20"} transition-all duration-300 ease-in-out relative min-h-[calc(100vh-3.5rem)] border-r border-gray-200 dark:border-gray-800`}>
      <div className="flex h-full flex-col gap-y-5 bg-white dark:bg-gray-900 px-6 pb-4">
        <div className="flex h-16 shrink-0 items-center">
          <button
            onClick={onToggle}
            className="absolute right-0 top-6 -mr-3 h-6 w-6 rounded-full bg-white dark:bg-gray-800 border border-gray-200 dark:border-gray-700 flex items-center justify-center"
          >
            {isExpanded ? (
              <ChevronLeftIcon className="h-4 w-4" />
            ) : (
              <ChevronRightIcon className="h-4 w-4" />
            )}
          </button>
        </div>
        <nav className="flex flex-1 flex-col">
          <ul role="list" className="flex flex-1 flex-col gap-y-7">
            <li>
              <ul role="list" className="-mx-2 space-y-1">
                {navigation.map((item) => {
                  const isActive = pathname === item.href;
                  return (
                    <li key={item.name}>
                      <Link
                        href={item.href}
                        className={`
                          group flex gap-x-3 rounded-md p-2 text-sm leading-6
                          ${isActive
                            ? "bg-gray-50 dark:bg-gray-800 text-blue-600 dark:text-blue-400"
                            : "text-gray-700 dark:text-gray-300 hover:text-blue-600 dark:hover:text-blue-400 hover:bg-gray-50 dark:hover:bg-gray-800"
                          }
                        `}
                      >
                        <item.icon className="h-6 w-6 shrink-0" aria-hidden="true" />
                        {isExpanded && <span className="whitespace-nowrap">{item.name}</span>}
                      </Link>
                    </li>
                  );
                })}
              </ul>
            </li>
          </ul>
        </nav>
      </div>
    </div>
  );
}