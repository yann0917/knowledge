package models

import (
	"github.com/yann0917/knowledge/base"
	"github.com/yann0917/knowledge/config"
)

type Topic struct {
	ID          int           `json:"id"`
	GroupID     int64         `json:"group_id"`
	TopicID     int64         `json:"topic_id"`
	Content     string        `json:"content"`
	Images      string        `json:"images"`
	ArticleID   string        `json:"article_id"`
	ArticleURL  string        `json:"article_url"`
	CreatedTime base.JSONTime `json:"created_time"`
	CreatedAt   base.JSONTime `json:"created_at"`
	UpdatedAt   base.JSONTime `json:"updated_at"`
	DeletedAt   base.JSONTime `json:"-"`
}

func (g *Topic) TableName() string {
	return "topic"
}

// FirstOrUpdate 创建或更新
func (g *Topic) FirstOrUpdate() (detail Topic, err error) {
	if err = config.DB.Where(&Topic{TopicID: g.TopicID}).
		Assign(&Topic{
			GroupID:     g.GroupID,
			TopicID:     g.TopicID,
			Content:     g.Content,
			Images:      g.Images,
			ArticleID:   g.ArticleID,
			ArticleURL:  g.ArticleURL,
			CreatedTime: g.CreatedTime,
		}).
		FirstOrCreate(&detail).Error; err != nil {
		return
	}
	return
}

func (g *Topic) List(list []Topic, err error) {
	db := config.DB.Order("created_at DESC")
	err = db.Find(&list).Error
	return
}
