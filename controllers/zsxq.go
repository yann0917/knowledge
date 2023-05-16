package controllers

import (
	"strings"
	"time"

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
	list, err := svc.GetGroupTopics(id, "digests", "20", "")
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
			if topic.ArticleURL != "" {
				article, err1 := svc.GetArticle(topic.ArticleURL)
				if err1 != nil {
					ctl.Error(err1)
					return
				}
				topic.RichContent = article
			}
		}
		if t.Type == "q&a" {
			topic.Content = t.Answer.Text + "\n\n" + t.Question.Text
		}
		topic.CreatedTime = base.JSONTime{Time: base.StrToTime(t.CreateTime)}

		topic.FirstOrUpdate()
	}
	ctl.Success(list)
}

func SyncColumns(c *gin.Context) {
	ctl := base.NewController(c)
	id := c.Param("id")
	list, err := svc.GetColumns(id)
	if err != nil {
		ctl.Error(err)
		return
	}

	// TODO: 总是报 1059 错误
	time.Sleep(1 * time.Second)

	for _, column := range list.Columns {
		topicList, err1 := svc.GetColumnTopics(id, base.Int642String(column.ColumnId))
		if err1 != nil {
			ctl.Error(err1)
			return
		}
		time.Sleep(1 * time.Second)

		for _, topic := range topicList.Topics {
			time.Sleep(3 * time.Second)
			info, err1 := svc.GetTopicInfo(base.Int642String(topic.TopicId))
			if err1 != nil {
				ctl.Error(err1)
				return
			}
			var topic models.Topic
			topic.GroupID = info.Topic.Group.GroupId
			topic.TopicID = info.Topic.TopicId
			if info.Topic.Type == "talk" {
				topic.Content = info.Topic.Talk.Text
				topic.ArticleID = info.Topic.Talk.Article.ArticleID
				topic.ArticleURL = info.Topic.Talk.Article.ArticleURL
				var imgs []string
				if len(info.Topic.Talk.Images) > 0 {
					for _, img := range info.Topic.Talk.Images {
						imgs = append(imgs, img.Large.Url)
					}
				}
				topic.Images = strings.Join(imgs, ",")
			}
			if info.Topic.Type == "q&a" {
				topic.Content = info.Topic.Answer.Text + "\n\n" + info.Topic.Question.Text
			}
			topic.CreatedTime = base.JSONTime{Time: base.StrToTime(info.Topic.CreateTime)}
			topic.FirstOrUpdate()
		}
	}

	ctl.Success(list)
}

func SyncColumnTopics(c *gin.Context) {
	ctl := base.NewController(c)
	gid := c.Param("id")
	cid := c.Param("column")
	list, err := svc.GetColumnTopics(gid, cid)
	if err != nil {
		ctl.Error(err)
		return
	}

	ctl.Success(list)
}

func SyncTopicInfo(c *gin.Context) {
	ctl := base.NewController(c)
	topicID := c.Param("id")
	list, err := svc.GetTopicInfo(topicID)
	if err != nil {
		ctl.Error(err)
		return
	}

	ctl.Success(list)
}

func GetArticle(c *gin.Context) {
	ctl := base.NewController(c)
	url := c.Query("url")
	list, err := svc.GetArticle(url)
	if err != nil {
		ctl.Error(err)
		return
	}
	ctl.Success(list)

}
