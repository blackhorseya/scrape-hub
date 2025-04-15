"use client";

import Link from "next/link";
import { MenuIcon, UserCircle, LogOut } from "lucide-react";
import { useUser } from "@auth0/nextjs-auth0/client";
import React from "react";

interface HeaderProps {
  onToggleSidebar: () => void;
}

export default function Header({ onToggleSidebar }: HeaderProps) {
  const { user, isLoading } = useUser();

  // 處理登入點擊事件
  const handleLoginClick = (e: React.MouseEvent) => {
    e.preventDefault();
    window.location.href = "/api/auth/login";
  };
  
  // 處理登出點擊事件
  const handleLogoutClick = (e: React.MouseEvent) => {
    e.preventDefault();
    window.location.href = "/api/auth/logout";
  };

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
          <div>
            {!isLoading && (
              user ? (
                <div className="flex items-center gap-4">
                  <Link
                    href="/profile"
                    className="text-sm text-gray-700 dark:text-gray-300 hover:text-blue-600 dark:hover:text-blue-400"
                  >
                    {user.name || user.email}
                  </Link>
                  <a
                    href="#"
                    onClick={handleLogoutClick}
                    className="inline-flex items-center gap-2 rounded-md bg-gray-100 dark:bg-gray-800 px-3 py-2 text-sm font-semibold text-gray-900 dark:text-gray-100 hover:bg-gray-200 dark:hover:bg-gray-700"
                  >
                    <LogOut className="h-5 w-5" aria-hidden="true" />
                    登出
                  </a>
                </div>
              ) : (
                <a
                  href="#"
                  onClick={handleLoginClick}
                  className="inline-flex items-center gap-2 rounded-md bg-blue-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-blue-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-blue-600"
                >
                  <UserCircle className="h-5 w-5" aria-hidden="true" />
                  登入
                </a>
              )
            )}
          </div>
        </div>
      </nav>
    </header>
  );
}