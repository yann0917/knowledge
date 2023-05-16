package controllers

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yann0917/knowledge/base"
	"github.com/yann0917/knowledge/models"
)

func UserSelf(c *gin.Context) {
	ctl := base.NewController(c)
	user, err := svc.GetUserSelf()
	if err != nil {
		ctl.Error(err)
		return
	}

	ctl.Success(user)
}

func SyncGroups(c *gin.Context) {
	ctl := base.NewController(c)
	list, err := svc.GetGroups()
	if err != nil {
		ctl.Error(err)
		return
	}
	for _, g := range list.Groups {
		var group models.Group
		group.GroupID = g.GroupId
		group.Name = g.Name
		group.FirstOrUpdate()
	}
	ctl.Success(list)
}

func SyncTopics(c *gin.Context) {
	ctl := base.NewController(c)
	id := c.Param("id")
	list, err := svc.GetTopics(id, "digests", "20", "")
	if err != nil {
		ctl.Error(err)
		return
	}

	for _, t := range list.Topics {
		var topic models.Topic
		topic.GroupID = t.Group.GroupId
		topic.TopicID = t.TopicId
		if t.Type == "talk" {
			topic.Content = t.Talk.Text
			topic.ArticleID = t.Talk.Article.ArticleID
			topic.ArticleURL = t.Talk.Article.ArticleURL
			var imgs []string
			if len(t.Talk.Images) > 0 {
				for _, img := range t.Talk.Images {
					imgs = append(imgs, img.Large.Url)
				}
			}
			topic.Images = strings.Join(imgs, ",")
		}
		if t.Type == "q&a" {
			topic.Content = t.Answer.Text + "\n\n" + t.Question.Text
		}
		topic.CreatedTime = base.JSONTime{Time: base.StrToTime(t.CreateTime)}

		topic.FirstOrUpdate()
	}
	ctl.Success(list)
}
