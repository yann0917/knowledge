package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yann0917/knowledge/base"
)

func SyncGroups(c *gin.Context) {
	ctl := base.NewController(c)
	svc.GetGroups()
	ctl.Success(1)
}

func SyncTopics(c *gin.Context) {
	ctl := base.NewController(c)

	ctl.Success(1)
}
