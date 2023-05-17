package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	. "github.com/yann0917/knowledge/controllers"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	gin.SetMode(gin.DebugMode)
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
		zsxq.GET("/:id/column/topic", GetColumnTopics)
	}

	return r
}
