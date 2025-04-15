package http

import (
	"fmt"

	"github.com/blackhorseya/scrape-hub/configs"
	"github.com/blackhorseya/scrape-hub/internal/delivery/middleware"
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
	
	// 建立 Auth0 中介層
	auth0Mid, err := middleware.NewAuth0Middleware(&cfg.Auth0)
	if err != nil {
		return nil, fmt.Errorf("初始化 Auth0 中介層失敗: %w", err)
	}

	// 註冊基礎路由
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// 受保護的路由群組
	protected := engine.Group("/api")
	protected.Use(auth0Mid.EnsureValidToken())

	return &serverImpl{
		engine: engine,
		cfg:    cfg,
	}, nil
}

// Engine 實作 Server 界面
func (s *serverImpl) Engine() *gin.Engine {
	return s.engine
}