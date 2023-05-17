package models

import (
	"time"

	"github.com/yann0917/knowledge/base"
	"github.com/yann0917/knowledge/config"
)

type Group struct {
	ID          int           `json:"id"`
	GroupID     int64         `json:"group_id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	CreatedTime time.Time     `json:"created_time"`
	CreatedAt   base.JSONTime `json:"created_at"`
	UpdatedAt   base.JSONTime `json:"updated_at"`
	DeletedAt   base.JSONTime `json:"-"`
}

func (g *Group) TableName() string {
	return "group"
}

// FirstOrUpdate 创建或更新
func (g *Group) FirstOrUpdate() (detail Group, err error) {
	if err = config.DB.Where(&Group{GroupID: g.GroupID}).
		Assign(&Group{
			GroupID:     g.GroupID,
			Name:        g.Name,
			Description: g.Description,
			CreatedTime: g.CreatedTime,
		}).
		FirstOrCreate(&detail).Error; err != nil {
		return
	}
	return
}

func (g *Group) List(list []Group, err error) {
	db := config.DB.Order("created_at DESC")
	err = db.Find(&list).Error
	return
}
