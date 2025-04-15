package http

import (
	"github.com/blackhorseya/scrape-hub/configs"
	"github.com/gin-gonic/gin"
)

// Server 定義 HTTP 伺服器界面
type Server interface {
	// Engine 回傳 Gin engine 實例
	Engine() *gin.Engine
}

// serverImpl 實作 HTTP 伺服器
type serverImpl struct {
	engine *gin.Engine
	cfg    *configs.Config
}

// NewServer 建立新的 HTTP 伺服器實例
func NewServer(cfg *configs.Config) (Server, error) {
	engine := gin.New()

	// 設定中介層
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())

	// 註冊基礎路由
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	return &serverImpl{
		engine: engine,
		cfg:    cfg,
	}, nil
}

// Engine 實作 Server 界面
func (s *serverImpl) Engine() *gin.Engine {
	return s.engine
}
