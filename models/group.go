package models

import (
	"github.com/yann0917/knowledge/base"
)

type Group struct {
	ID        int           `json:"id"`
	GroupID   int64         `json:"group_id"`
	Name      string        `json:"name"`
	CreatedAt base.JSONTime `json:"created_at"`
	UpdatedAt base.JSONTime `json:"updated_at"`
	DeletedAt base.JSONTime `json:"-"`
}
