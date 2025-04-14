package configs

import (
	"fmt"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

// ServerConfig 定義伺服器設定
type ServerConfig struct {
	Host string `env:"SERVER_HOST" envDefault:"localhost"`
	Port int    `env:"SERVER_PORT" envDefault:"8080"`
}

// Config 定義應用程式設定
type Config struct {
	Server ServerConfig
}

// LoadFromEnv 從環境變數載入設定
// 它會先嘗試從專案根目錄的 .env 檔案載入變數，然後解析環境變數到 Config 物件
func LoadFromEnv() (*Config, error) {
	// 嘗試載入 .env 檔案
	if err := godotenv.Load(); err != nil {
		// 若 .env 檔案不存在，僅記錄訊息但不中斷執行
		fmt.Printf("警告: 無法載入 .env 檔案: %v\n", err)
	}

	cfg := &Config{}
	opts := env.Options{
		Prefix: "",
	}

	if err := env.ParseWithOptions(cfg, opts); err != nil {
		return nil, fmt.Errorf("解析環境變數失敗: %w", err)
	}

	return cfg, nil
}
