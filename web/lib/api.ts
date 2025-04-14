// API 錯誤類型
export class ApiError extends Error {
  constructor(public status: number, message: string) {
    super(message);
    this.name = 'ApiError';
  }
}

// API 請求設定介面
interface RequestOptions extends RequestInit {
  accessToken?: string;
}

// API 請求工具函式
export async function fetchApi<T>(
  endpoint: string,
  options: RequestOptions = {}
): Promise<T> {
  const apiUrl = process.env.NEXT_PUBLIC_API_URL;
  if (!apiUrl) {
    throw new Error('API URL 未設定');
  }

  const url = `${apiUrl}${endpoint}`;
  
  // 建立 headers 物件
  const headers: Record<string, string> = {
    'Content-Type': 'application/json',
  };

  // 如果有 access token，加入到 headers
  if (options.accessToken) {
    headers['Authorization'] = `Bearer ${options.accessToken}`;
  }

  const response = await fetch(url, {
    ...options,
    headers: {
      ...headers,
      ...(options.headers as Record<string, string>),
    },
  });

  // 處理非 2xx 的回應
  if (!response.ok) {
    throw new ApiError(
      response.status,
      response.statusText || '請求失敗'
    );
  }

  // 解析並回傳回應內容
  return response.json();
}