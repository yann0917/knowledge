package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	. "github.com/yann0917/knowledge/controllers"
	"github.com/yann0917/knowledge/middlewares"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery(), middlewares.Cors())

	gin.SetMode(gin.DebugMode)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pang"})
	})

	zsxq := r.Group("/zsxq")
	{
		zsxq.GET("/user/self", UserSelf)
		zsxq.GET("/sync/group", SyncGroups)
		zsxq.GET("/sync/:id/topic", SyncTopics)
		zsxq.GET("/sync/:id/column", SyncColumns)
		zsxq.GET("/article", GetArticle)
		zsxq.GET("/:id/:scope/topic", GetTopics)
	}

	return r
}
