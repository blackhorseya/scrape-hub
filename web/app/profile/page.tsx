"use client";

import { useUser } from "@auth0/nextjs-auth0/client";
import { Copy } from "lucide-react";
import { useState, useEffect } from "react";

export default function ProfilePage() {
  const { user, isLoading } = useUser();
  const [copied, setCopied] = useState<string>("");
  const [accessToken, setAccessToken] = useState<string>("");
  const [tokenError, setTokenError] = useState<string>("");

  useEffect(() => {
    const fetchToken = async () => {
      try {
        const response = await fetch('/api/auth/token');
        const data = await response.json();
        if (data.error) {
          setTokenError(data.error);
        } else {
          setAccessToken(data.accessToken);
        }
      } catch (error) {
        setTokenError('無法取得 access token');
        console.error('Error fetching access token:', error);
      }
    };

    if (user) {
      fetchToken();
    }
  }, [user]);

  if (isLoading) {
    return (
      <div className="flex items-center justify-center min-h-[50vh]">
        <div className="animate-spin rounded-full h-8 w-8 border-t-2 border-b-2 border-blue-500"></div>
      </div>
    );
  }

  if (!user) {
    return (
      <div className="text-center py-12">
        <h1 className="text-2xl font-semibold text-gray-900 dark:text-gray-100">
          請先登入以查看個人資料
        </h1>
      </div>
    );
  }

  const handleCopy = async (text: string, type: string) => {
    try {
      await navigator.clipboard.writeText(text);
      setCopied(type);
      setTimeout(() => setCopied(""), 2000);
    } catch (err) {
      console.error("無法複製文字:", err);
    }
  };

  return (
    <div className="max-w-3xl mx-auto py-8 px-4 sm:px-6 lg:px-8">
      <div className="bg-white dark:bg-gray-800 shadow rounded-lg">
        <div className="px-4 py-5 sm:p-6">
          <h1 className="text-2xl font-semibold text-gray-900 dark:text-gray-100 mb-8">
            個人資料
          </h1>

          <div className="space-y-6">
            {user.picture && (
              <div className="flex items-center gap-4">
                <img
                  src={user.picture}
                  alt={user.name || "使用者頭像"}
                  className="h-20 w-20 rounded-full"
                />
              </div>
            )}

            <div>
              <h3 className="text-sm font-medium text-gray-500 dark:text-gray-400">
                名稱
              </h3>
              <p className="mt-1 text-lg text-gray-900 dark:text-gray-100">
                {user.name || "未設定"}
              </p>
            </div>

            <div>
              <h3 className="text-sm font-medium text-gray-500 dark:text-gray-400">
                電子郵件
              </h3>
              <p className="mt-1 text-lg text-gray-900 dark:text-gray-100">
                {user.email}
              </p>
            </div>

            <div>
              <h3 className="text-sm font-medium text-gray-500 dark:text-gray-400">
                使用者 ID
              </h3>
              <div className="mt-1 flex items-center gap-2">
                <code className="px-2 py-1 text-sm bg-gray-100 dark:bg-gray-700 rounded">
                  {user.sub}
                </code>
                <button
                  onClick={() => handleCopy(user.sub || "", "id")}
                  className="inline-flex items-center gap-1 px-2 py-1 text-sm text-blue-600 dark:text-blue-400 hover:text-blue-500"
                  title="複製使用者 ID"
                >
                  <Copy className="h-4 w-4" />
                  {copied === "id" ? "已複製！" : "複製"}
                </button>
              </div>
            </div>

            <div>
              <h3 className="text-sm font-medium text-gray-500 dark:text-gray-400">
                Access Token
              </h3>
              <div className="mt-1">
                {tokenError ? (
                  <p className="text-red-500 dark:text-red-400">{tokenError}</p>
                ) : accessToken ? (
                  <div className="flex items-start gap-2">
                    <code className="px-2 py-1 text-sm bg-gray-100 dark:bg-gray-700 rounded break-all max-w-2xl">
                      {accessToken}
                    </code>
                    <button
                      onClick={() => handleCopy(accessToken, "token")}
                      className="inline-flex items-center gap-1 px-2 py-1 text-sm text-blue-600 dark:text-blue-400 hover:text-blue-500 shrink-0"
                      title="複製 Access Token"
                    >
                      <Copy className="h-4 w-4" />
                      {copied === "token" ? "已複製！" : "複製"}
                    </button>
                  </div>
                ) : (
                  <div className="animate-pulse h-6 w-64 bg-gray-200 dark:bg-gray-700 rounded"></div>
                )}
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}