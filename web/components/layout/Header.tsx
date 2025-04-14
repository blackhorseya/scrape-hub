"use client";

import Link from "next/link";
import { MenuIcon } from "lucide-react";
import React from "react";

interface HeaderProps {
  onToggleSidebar: () => void;
}

export default function Header({ onToggleSidebar }: HeaderProps) {
  return (
    <header className="bg-white shadow-sm dark:bg-gray-900 sticky top-0 z-50">
      <nav className="mx-auto px-4 sm:px-6 lg:px-8" aria-label="頂部導覽">
        <div className="flex h-14 items-center justify-between">
          <div className="flex items-center">
            <button
              type="button"
              className="lg:hidden -m-2.5 inline-flex items-center justify-center rounded-md p-2.5 text-gray-700 dark:text-gray-200"
              onClick={onToggleSidebar}
            >
              <span className="sr-only">開啟側邊選單</span>
              <MenuIcon className="h-6 w-6" aria-hidden="true" />
            </button>
            <div className="ml-4 flex lg:ml-0">
              <Link href="/" className="text-xl font-semibold">
                Scrape Hub
              </Link>
            </div>
          </div>
          <div className="ml-10 flex items-center space-x-4">
            <Link
              href="/tasks"
              className="text-sm font-medium text-gray-700 dark:text-gray-200 hover:text-gray-900 dark:hover:text-white"
            >
              任務清單
            </Link>
            <Link
              href="/settings"
              className="text-sm font-medium text-gray-700 dark:text-gray-200 hover:text-gray-900 dark:hover:text-white"
            >
              設定
            </Link>
          </div>
        </div>
      </nav>
    </header>
  );
}