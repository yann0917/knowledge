package models

import (
	"net/url"
	"regexp"
	"strings"
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

func (g *Topic) List(filterTag bool) (list []Topic, err error) {
	db := config.DB.Order("create_time DESC").Where(g)
	err = db.Find(&list).Error
	if filterTag {
		list = TopicsReplaceTag(list)
	}
	return
}

// TopicsReplaceTag replace mention, hashtag, web tag
func TopicsReplaceTag(list []Topic) (topics []Topic) {
	for i, topic := range list {
		Content := topic.Content
		if topic.Content != "" {
			// replace mention, hashtag tag
			reg := `<e type="mention"[^>]*>.*?|<e type="mention".*? \/>`
			matches := regexp.MustCompile(reg).FindAllString(Content, -1)
			for _, match := range matches {
				Content = strings.ReplaceAll(Content, match, "")
			}

			reg = `<e type="hashtag"[^>]*>.*?|<e type="hashtag".*? \/>`
			matches = regexp.MustCompile(reg).FindAllString(Content, -1)
			for _, match := range matches {
				Content = strings.ReplaceAll(Content, match, "")
			}
		}

		// replace web tag with a tag
		if Content != "" {
			reg := `<e type="web"[^>]*>.*?|<e type="web".*? \/>`
			matches := regexp.MustCompile(reg).FindAllString(Content, -1)
			for _, match := range matches {
				reg2 := `href="(.*?)" title="(.*?)"`
				matches2 := regexp.MustCompile(reg2).FindAllStringSubmatch(match, -1)
				if len(matches2) > 0 {
					result := matches2[0]
					if len(result) == 3 {
						href, _ := url.PathUnescape(result[1])
						title, _ := url.PathUnescape(result[2])
						newText := "<a href='" + href + "' target='_black'>" + title + "</a>"
						Content = strings.ReplaceAll(Content, match, newText)
					}
				}
			}
		}
		list[i].Content = Content
	}
	topics = list
	return
}

func (g *Topic) GetLast() (detail Topic, err error) {
	db := config.DB.Where(g).Order("create_time asc")
	err = db.
		First(&detail).
		Error
	return
}

func (g *Topic) GetLatest() (detail Topic, err error) {
	db := config.DB.Where(g).Order("create_time desc")
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

func ConvertToMd(list []Topic) (res string) {
	for _, topic := range list {
		Content := topic.Content
		if topic.Title != "" {
			res += base.GetMdHeader(2) + topic.Title + "\r\n\r\n"
		} else {
			if Content != "" {
				cont := []rune(Content)
				res += base.GetMdHeader(2) + string(cont[0:10]) + "\r\n\r\n"
			} else {
				res += base.GetMdHeader(2) + base.Int642String(topic.TopicID) + "\r\n\r\n"
			}
		}
		res += "> 创建时间: " + topic.CreateTime.Format("2006-01-02 15:04:05") + "\r\n\r\n"
		res += Content + "\r\n\r\n"
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
		res += `<div style="page-break-after: always;"></div>` + "\r\n\r\n"
	}
	return
}

func ConvertToHtml(list []Topic) (res string) {
	res = base.GenStartHtml()
	for _, topic := range list {
		Content := topic.Content
		if topic.Title != "" {
			res += base.GenHLevelHtml(0, true) + topic.Title + base.GenHLevelHtml(0, false)
		} else {
			if Content != "" {
				cont := []rune(Content)
				res += base.GenHLevelHtml(0, true) + string(cont[0:10]) + base.GenHLevelHtml(0, false)
			} else {
				res += base.GenHLevelHtml(0, true) + base.Int642String(topic.TopicID) + base.GenHLevelHtml(0, false)
			}
		}
		res += "<p> 创建时间: " + topic.CreateTime.Format("2006-01-02 15:04:05") + "</p>\r\n\r\n"
		res += Content
		if topic.Images != "" {
			imgs := strings.Split(topic.Images, ",")
			for _, img := range imgs {
				res += "<img src='" + img + "' alt='' />"
			}
		}

		if topic.RichContent != "" {
			res += "\r\n"
			res += base.GenHLevelHtml(1, true) + "文章" + base.GenHLevelHtml(1, false)
			res += topic.RichContent + "\r\n"
		}
		res += base.GenPageBreak()
	}
	res += base.GenEndHtml()
	return
}
