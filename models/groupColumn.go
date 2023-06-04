package models

import (
	"time"

	"github.com/yann0917/knowledge/base"
	"github.com/yann0917/knowledge/config"
	"gorm.io/gorm"
)

type GroupColumn struct {
	ID                  int           `json:"id"`
	GroupID             int64         `json:"group_id"`
	ColumnID            int64         `json:"column_id"`
	Name                string        `json:"name"`
	CoverURL            string        `json:"cover_url"`
	CreateTime          time.Time     `json:"create_time"`
	LastTopicAttachTime time.Time     `json:"last_topic_attach_time"`
	CreatedAt           base.JSONTime `json:"created_at"`
	UpdatedAt           base.JSONTime `json:"updated_at"`
	DeletedAt           base.JSONTime `json:"-"`

	Topics []Topic `json:"topics" gorm:"many2many:column_topic;foreignKey:ColumnID;joinForeignKey:ColumnID;References:TopicID;joinReferences:TopicID"`
}

func (g *GroupColumn) TableName() string {
	return "group_column"
}

// FirstOrUpdate 创建或更新
func (g *GroupColumn) FirstOrUpdate() (detail GroupColumn, err error) {
	if err = config.DB.Where(&GroupColumn{ColumnID: g.ColumnID}).
		Assign(&GroupColumn{
			GroupID:             g.GroupID,
			ColumnID:            g.ColumnID,
			Name:                g.Name,
			CoverURL:            g.CoverURL,
			CreateTime:          g.CreateTime,
			LastTopicAttachTime: g.LastTopicAttachTime,
		}).
		FirstOrCreate(&detail).Error; err != nil {
		return
	}
	return
}

func (g *GroupColumn) List(filterTag bool) (list []GroupColumn, err error) {
	db := config.DB.Order("created_at ASC").Where(g).
		Preload("Topics", func(db *gorm.DB) *gorm.DB {
			return db.Not("content", "").Order("topic.create_time DESC")
		})
	err = db.Find(&list).Error
	if filterTag {
		for i, column := range list {
			list[i].Topics = TopicsReplaceTag(column.Topics)
		}
	}
	return
}

func (g *GroupColumn) ConvertToMd(list []GroupColumn) (res string) {
	for _, column := range list {
		res += base.GetMdHeader(1) + column.Name + "\r\n\r\n"
		res += ConvertToMd(column.Topics)
	}
	return
}
