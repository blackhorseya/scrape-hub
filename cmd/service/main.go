package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/blackhorseya/scrape-hub/configs"
)

func main() {
	// 載入設定
	cfg, err := configs.LoadFromEnv()
	if err != nil {
		log.Fatalf("載入設定失敗: %v", err)
	}

	fmt.Printf("伺服器將在 %s:%d 啟動\n", cfg.Server.Host, cfg.Server.Port)

	// 設定信號處理
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	log.Println("服務啟動中...")

	// 等待終止信號
	<-signals
	log.Println("服務關閉中...")
}
