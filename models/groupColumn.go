package models

import (
	"net/url"
	"regexp"
	"strings"
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

func (g *GroupColumn) List() (list []GroupColumn, err error) {
	db := config.DB.Order("created_at ASC").Where(g).
		Preload("Topics", func(db *gorm.DB) *gorm.DB {
			return db.Order("topic.create_time DESC")
		})
	err = db.Find(&list).Error
	return
}

func (g *GroupColumn) ConvertToMd(list []GroupColumn) (res string) {
	for _, column := range list {
		res += base.GetMdHeader(1) + column.Name + "\r\n\r\n"

		for _, topic := range column.Topics {
			if topic.Title != "" {
				res += base.GetMdHeader(2) + topic.Title + "\r\n\r\n"
			} else {
				if topic.Content != "" {
					cont := []rune(topic.Content)
					res += base.GetMdHeader(2) + string(cont[0:10]) + "\r\n\r\n"
				} else {
					res += base.GetMdHeader(2) + base.Int642String(topic.TopicID) + "\r\n\r\n"
				}
			}

			res += "> 创建时间: " + base.TimeToStr(topic.CreateTime) + "\r\n\r\n"

			if topic.Content != "" {
				reg := `<e type="web"[^>]*>.*?|<.*? \/>`
				matches := regexp.MustCompile(reg).FindAllString(topic.Content, -1)
				Content := topic.Content
				for _, match := range matches {
					reg2 := `href="(.*?)" title="(.*?)"`
					matches2 := regexp.MustCompile(reg2).FindAllStringSubmatch(match, -1)
					if len(matches2) > 0 {
						result := matches2[0]
						if len(result) == 3 {
							href, _ := url.PathUnescape(result[1])
							title, _ := url.PathUnescape(result[2])
							newText := "[" + title + "](" + href + ")"
							Content = strings.ReplaceAll(Content, match, newText)
						}
					}
				}
				res += Content + "\r\n\r\n"
			}

			if topic.Images != "" {
				imgs := strings.Split(topic.Images, ",")
				for _, img := range imgs {
					res += "![](" + img + ")\r\n"
				}
			}

			if topic.RichContent != "" {
				res += "---\r\n\r\n"
				res += base.GetMdHeader(3) + "文章" + "\r\n\r\n"
				res += topic.RichContent + "\r\n\r\n"
			}
		}
	}
	return
}
