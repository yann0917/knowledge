package models

import (
	"time"

	"github.com/yann0917/knowledge/base"
	"github.com/yann0917/knowledge/config"
)

type Topic struct {
	ID          int           `json:"id"`
	GroupID     int64         `json:"group_id"`
	TopicID     int64         `json:"topic_id"`
	Title       string        `json:"title"`
	Content     string        `json:"content"`
	RichContent string        `json:"rich_content"`
	Images      string        `json:"images"`
	IsDigests   bool          `json:"is_digests" gorm:"default:0"`
	ArticleID   string        `json:"article_id"`
	ArticleURL  string        `json:"article_url"`
	CreateTime  time.Time     `json:"create_time"`
	CreatedAt   base.JSONTime `json:"created_at"`
	UpdatedAt   base.JSONTime `json:"updated_at"`
	DeletedAt   base.JSONTime `json:"-"`
}

func (g *Topic) TableName() string {
	return "topic"
}

type ColumnTopic struct {
	ID                   int           `json:"id"`
	ColumnID             int64         `json:"column_id"`
	TopicID              int64         `json:"topic_id"`
	AttachedToColumnTime time.Time     `json:"attached_to_column_time"`
	CreatedAt            base.JSONTime `json:"created_at"`
	UpdatedAt            base.JSONTime `json:"updated_at"`
	DeletedAt            base.JSONTime `json:"-"`
}

func (g *ColumnTopic) TableName() string {
	return "column_topic"
}

// FirstOrUpdate 创建或更新
func (g *Topic) FirstOrUpdate() (detail Topic, err error) {
	if err = config.DB.Where(&Topic{TopicID: g.TopicID}).
		Assign(&Topic{
			GroupID:     g.GroupID,
			TopicID:     g.TopicID,
			Title:       g.Title,
			Content:     g.Content,
			RichContent: g.RichContent,
			Images:      g.Images,
			IsDigests:   g.IsDigests,
			ArticleID:   g.ArticleID,
			ArticleURL:  g.ArticleURL,
			CreateTime:  g.CreateTime,
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

func (g *Topic) GetLast() (detail Topic, err error) {
	db := config.DB.Where(g).Order("created_time asc")
	err = db.
		First(&detail).
		Error
	return
}

func (g *Topic) GetLatest() (detail Topic, err error) {
	db := config.DB.Where(g).Order("created_time desc")
	err = db.
		First(&detail).
		Error
	return
}

// FirstOrUpdate 创建或更新
func (g *ColumnTopic) FirstOrUpdate() (detail ColumnTopic, err error) {
	if err = config.DB.Where(&ColumnTopic{TopicID: g.TopicID}).
		Assign(&ColumnTopic{
			TopicID:              g.TopicID,
			ColumnID:             g.ColumnID,
			AttachedToColumnTime: g.AttachedToColumnTime,
		}).
		FirstOrCreate(&detail).Error; err != nil {
		return
	}
	return
}
