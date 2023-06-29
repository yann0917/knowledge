package controllers

import (
	"bytes"
	"errors"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yann0917/knowledge/base"
	"github.com/yann0917/knowledge/models"
	"github.com/yann0917/knowledge/service"
	"gorm.io/gorm"
)

const PageSize = 20   //
const PullDelay = 300 // 延迟 ms
const UnixTime = "2006-01-02T15:04:05.000+0800"

func UserSelf(c *gin.Context) {
	ctl := base.NewController(c)
	user, err := svc.GetUserSelf()
	if err != nil {
		ctl.Error(err)
		return
	}

	ctl.Success(user)
}

// SyncGroups 同步加入的群组
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

// SyncTopics 同步群组所有主题
func SyncTopics(c *gin.Context) {
	ctl := base.NewController(c)
	id := c.Param("id")
	var topic models.Topic
	topic.GroupID, _ = base.String2Int64(id)
	latestTopic, err := topic.GetLatest()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		ctl.Error(err)
		return
	}
	if !latestTopic.CreateTime.IsZero() {
		endTime := latestTopic.CreateTime.Add(-1 * time.Second)
		pullTopicsByPeriod(id, endTime, base.StrToTimeUtc8(UnixTime))
		time.Sleep(5 * time.Second)

		lastTopic, err1 := topic.GetLast()
		if err1 != nil && !errors.Is(err1, gorm.ErrRecordNotFound) {
			ctl.Error(err1)
			return
		}
		// 拉取之后的
		endTime = lastTopic.CreateTime.Add(-1 * time.Second)
		pullTopicsByPeriod(id, endTime, base.StrToTimeUtc8(UnixTime))

		time.Sleep(5 * time.Second)
		// 拉取最新
		pullTopicsByPeriod(id, time.Now(), endTime)

	} else {
		// 全量拉取
		pullTopicsByPeriod(id, time.Now(), base.StrToTimeUtc8(UnixTime))
	}

	ctl.Success(1)
}

func pullTopicsByPeriod(groupID string, startTime, endTime time.Time) {
	waitGroup := sync.WaitGroup{}
	returnSize := PageSize
	for {
		if returnSize < PageSize || startTime.Before(endTime) {
			break
		}

		list, err1 := svc.GetGroupTopics(groupID, "all", "20", base.TimeToStr(startTime))
		if err1 != nil {
			return
		}

		returnSize = len(list.Topics)
		if returnSize == 0 {
			return
		}
		startTime = base.StrToTimeUtc8(list.Topics[returnSize-1].CreateTime)

		waitGroup.Add(1)
		go saveTopicsByPage(&waitGroup, list)
		time.Sleep(PullDelay * time.Millisecond)
	}
	waitGroup.Wait()
}

func saveTopicsByPage(w *sync.WaitGroup, list service.Topics) {

	for _, t := range list.Topics {
		var topic models.Topic
		topic.GroupID = t.Group.GroupId
		topic.TopicID = t.TopicId
		topic.Title = t.Title
		topic.IsDigests = t.Digested

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
					return
				}
				topic.RichContent = article
			}
		}
		if t.Type == "q&a" {
			topic.Content = t.Answer.Text + "|" + t.Question.Text
		}
		topic.CreateTime = base.StrToTimeUtc8(t.CreateTime)

		topic.FirstOrUpdate()
	}
	w.Done()
}

// SyncColumns 同步群组分类主题
func SyncColumns(c *gin.Context) {
	ctl := base.NewController(c)
	id := c.Param("id")
	list, err := svc.GetColumns(id)
	if err != nil {
		ctl.Error(err)
		return
	}

	for _, column := range list.Columns {
		var gc models.GroupColumn
		gc.GroupID, _ = base.String2Int64(id)
		gc.ColumnID = column.ColumnId
		gc.Name = column.Name
		gc.CreateTime = base.StrToTimeUtc8(column.CreateTime)
		gc.LastTopicAttachTime = base.StrToTimeUtc8(column.LastTopicAttachTime)
		gc.CoverURL = column.CoverUrl
		gc.FirstOrUpdate()
	}

	for _, column := range list.Columns {
		topicList, err1 := svc.GetColumnTopics(id, base.Int642String(column.ColumnId))
		if err1 != nil {
			ctl.Error(err1)
			return
		}
		// 不要优化，否则会触发反爬虫机制报 1059 错误
		for _, topic := range topicList.Topics {
			var ct models.ColumnTopic
			ct.ColumnID = column.ColumnId
			ct.TopicID = topic.TopicId
			ct.AttachedToColumnTime = base.StrToTimeUtc8(topic.AttachedToColumnTime)
			ct.FirstOrUpdate()
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

func GetTopics(c *gin.Context) {
	ctl := base.NewController(c)
	id := c.Param("id")
	scope := c.Param("scope")

	// 1-md, 2-pdf
	dtype := c.DefaultQuery("type", "1")

	groupID, _ := base.String2Int64(id)
	var g models.Group
	g.GroupID = groupID
	detail, err := g.Detail()
	if err != nil {
		ctl.Error(err)
		return
	}

	switch scope {
	case "column":
		title := detail.Name + "-导读分类"
		var gc models.GroupColumn
		gc.GroupID = groupID
		list, err1 := gc.List(true)
		if err1 != nil {
			ctl.Error(err1)
			return
		}

		switch dtype {
		case "1":
			res := gc.ConvertToMd(list)
			name := base.FileName(title, "md")
			err = base.SaveFile(name, res)
			if err != nil {
				ctl.Error(err)
				return
			}
		case "2":
			buf := new(bytes.Buffer)
			res := base.GenStartHtml()
			res += gc.ConvertToHtml(list)
			res += base.GenEndHtml()
			// name := base.FileName(title, "html")
			// err = base.SaveFile(name, res)
			buf.Write([]byte(res))
			path := base.GetCurrentDirectory()
			filePreName := filepath.Join(path, base.FileName(title, ""))
			fileName, err := base.FilePath(filePreName, "pdf", false)
			if err != nil {
				ctl.Error(err)
				return
			}
			err = base.GenPdf(buf, fileName)
			if err != nil {
				ctl.Error(err)
				return
			}
		}

		ctl.Success(list)
	case "digest":

		title := detail.Name + "-精华主题"
		var t models.Topic
		t.GroupID = groupID
		t.IsDigests = true

		list, err1 := t.List(true)
		if err1 != nil {
			ctl.Error(err1)
			return
		}

		switch dtype {

		case "1":
			res := models.ConvertToMd(list)
			name := base.FileName(title, "md")
			err = base.SaveFile(name, res)
			if err != nil {
				ctl.Error(err)
				return
			}
		case "2":
			buf := new(bytes.Buffer)
			res := base.GenStartHtml()
			res += models.ConvertToHtml(list)
			res += base.GenEndHtml()
			// name := base.FileName(title, "html")
			// err = base.SaveFile(name, res)
			buf.Write([]byte(res))
			path := base.GetCurrentDirectory()
			filePreName := filepath.Join(path, base.FileName(title, ""))
			fileName, err := base.FilePath(filePreName, "pdf", false)
			if err != nil {
				ctl.Error(err)
				return
			}
			err = base.GenPdf(buf, fileName)
			if err != nil {
				ctl.Error(err)
				return
			}
		}
		ctl.Success(list)
	}

}
