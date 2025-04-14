package middleware

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/blackhorseya/scrape-hub/configs"
	"github.com/gin-gonic/gin"
)

// CustomClaims 包含自訂的 JWT claims
type CustomClaims struct {
	Scope string `json:"scope"`
}

// Validate 實作 validator.CustomClaims 界面
func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

// Auth0Middleware 提供 Auth0 JWT 驗證的中介層結構
type Auth0Middleware struct {
	validator *validator.Validator
}

// NewAuth0Middleware 建立新的 Auth0 中介層實例
func NewAuth0Middleware(cfg *configs.Auth0Config) (*Auth0Middleware, error) {
	issuerURL, err := url.Parse("https://" + cfg.Domain + "/")
	if err != nil {
		return nil, err
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{cfg.Audience},
		validator.WithCustomClaims(
			func() validator.CustomClaims {
				return &CustomClaims{}
			},
		),
	)
	if err != nil {
		return nil, err
	}

	return &Auth0Middleware{
		validator: jwtValidator,
	}, nil
}

// EnsureValidToken 驗證 JWT token 的有效性
func (m *Auth0Middleware) EnsureValidToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := extractToken(c.Request)
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "未提供 token",
			})
			return
		}

		claims, err := m.validator.ValidateToken(c.Request.Context(), token)
		if err != nil {
			log.Printf("token 驗證失敗: %v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "token 驗證失敗",
			})
			return
		}

		// 將驗證後的 claims 注入到 context 中
		c.Set("claims", claims)
		c.Next()
	}
}

// extractToken 從請求標頭中提取 bearer token
func extractToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}
