package router

import (
	"net/http"

	"github.com/dragonforce2010/chatgpt-service/domain/chat"
	"github.com/gin-gonic/gin"
)

type Router struct {
	router     *gin.Engine
	chatRouter *chat.ChatRouter
}

func NewRouter(chatRouter *chat.ChatRouter, chatHandler *chat.ChatHandler) *Router {
	return &Router{
		chatRouter: chatRouter,
	}
}

// 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

func (r *Router) InitRoutes() {
	r.router = gin.Default()
	r.router.Use(Cors())
	//r.router.Use(cors.New(cors.Config{
	//	AllowAllOrigins:        true,
	//	AllowOrigins: 	 		[]string{"*"},
	//	AllowMethods:           []string{"POST", "GET","PUT","DELETE","OPTIONS"},
	//	AllowHeaders:           []string{"Origin", "Content-Length", "Content-Type", "Access-Control-Allow-Headers", "Authorization", "X-Requested-With"},
	//	AllowCredentials:       true,
	//}))
	r.chatRouter.Init(r.router)
	r.router.GET("/api/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}

func (r *Router) Run(address string) {
	r.InitRoutes()
	//// # Headers
	// Allow CORS
	r.router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	})
	r.router.Run(address)
}
