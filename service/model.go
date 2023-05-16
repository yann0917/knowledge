package service

type Summary struct {
	HasColumns bool   `json:"has_columns"`
	Title      string `json:"title"`
}

type Menu struct {
	MenuId                int64  `json:"menu_id"`
	Title                 string `json:"title"`
	Preset                bool   `json:"preset"`
	PresetType            string `json:"preset_type"`
	LatestTopicCreateTime string `json:"latest_topic_create_time,omitempty"`
}

type Menus struct {
	Menus []Menu `json:"menus"`
}

type Column struct {
	ColumnId   int64  `json:"column_id"`
	Name       string `json:"name"`
	CoverUrl   string `json:"cover_url"`
	Statistics struct {
		TopicsCount int `json:"topics_count"`
	} `json:"statistics"`
	CreateTime          string `json:"create_time"`
	LastTopicAttachTime string `json:"last_topic_attach_time"`
}

type Columns struct {
	Columns []Column `json:"columns"`
}

type Group struct {
	GroupId int64  `json:"group_id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
}

type User struct {
	UserId      int64  `json:"user_id"`
	Name        string `json:"name"`
	AvatarUrl   string `json:"avatar_url"`
	Description string `json:"description,omitempty"`
	Location    string `json:"location"`
	Alias       string `json:"alias,omitempty"`
	Number      int    `json:"number,omitempty"`
	UniqueId    string `json:"unique_id,omitempty"`
}

type T2 struct {
	User User `json:"user"`
	Chat struct {
		Identifier string `json:"identifier"`
	} `json:"chat"`
	Accounts struct {
		Phone struct {
			CountryCode string `json:"country_code"`
			PhoneNumber string `json:"phone_number"`
		} `json:"phone"`
		Wechat struct {
			Name      string `json:"name"`
			AvatarUrl string `json:"avatar_url"`
		} `json:"wechat"`
	} `json:"accounts"`
	IdentityStatus   string `json:"identity_status"`
	SubscribedWechat bool   `json:"subscribed_wechat"`
	Subscriptions    struct {
		SubscribedXiaomiquanvip bool `json:"subscribed_xiaomiquanvip"`
		SubscribedXingqiusvip   bool `json:"subscribed_xingqiusvip"`
	} `json:"subscriptions"`
	AssociatedEnterprise bool `json:"associated_enterprise"`
}

type Article struct {
	Title            string `json:"title"`
	ArticleId        string `json:"article_id"`
	ArticleUrl       string `json:"article_url"`
	InlineArticleUrl string `json:"inline_article_url"`
}

type ImageMeta struct {
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Size   int    `json:"size,omitempty"`
}

type Image struct {
	ImageId   int64     `json:"image_id"`
	Type      string    `json:"type"`
	Thumbnail ImageMeta `json:"thumbnail"`
	Large     ImageMeta `json:"large"`
	Original  ImageMeta `json:"original,omitempty"`
}

type File struct {
	FileId        int64  `json:"file_id"`
	Name          string `json:"name"`
	Hash          string `json:"hash"`
	Size          int    `json:"size"`
	DownloadCount int    `json:"download_count"`
	CreateTime    string `json:"create_time"`
	Duration      int    `json:"duration"`
}

type Topics struct {
	Topics []struct {
		TopicId int64  `json:"topic_id"`
		Group   Group  `json:"group"`
		Type    string `json:"type"`
		Talk    struct {
			Owner   User    `json:"owner"`
			Text    string  `json:"text"`
			Article Article `json:"article,omitempty"`
			Images  []Image `json:"images,omitempty"`
			Files   []File  `json:"files,omitempty"`
		} `json:"talk"`
		LatestLikes []struct {
			CreateTime string `json:"create_time"`
			Owner      User   `json:"owner"`
		} `json:"latest_likes"`
		ShowComments []struct {
			CommentId       int64   `json:"comment_id"`
			CreateTime      string  `json:"create_time"`
			Owner           User    `json:"owner"`
			Text            string  `json:"text"`
			LikesCount      int     `json:"likes_count"`
			RewardsCount    int     `json:"rewards_count"`
			Sticky          bool    `json:"sticky"`
			RepliesCount    int     `json:"replies_count,omitempty"`
			ParentCommentId int64   `json:"parent_comment_id,omitempty"`
			Repliee         User    `json:"repliee,omitempty"`
			Images          []Image `json:"images,omitempty"`
		} `json:"show_comments,omitempty"`
		LikesCount    int    `json:"likes_count"`
		RewardsCount  int    `json:"rewards_count"`
		CommentsCount int    `json:"comments_count"`
		ReadingCount  int    `json:"reading_count"`
		ReadersCount  int    `json:"readers_count"`
		Digested      bool   `json:"digested"`
		Sticky        bool   `json:"sticky"`
		CreateTime    string `json:"create_time"`
		ModifyTime    string `json:"modify_time,omitempty"`
		UserSpecific  struct {
			Liked      bool `json:"liked"`
			Subscribed bool `json:"subscribed"`
		} `json:"user_specific"`
		Title   string `json:"title,omitempty"`
		Columns []struct {
			ColumnId int64  `json:"column_id"`
			Name     string `json:"name"`
		} `json:"columns"`
	} `json:"topics"`
}

type Topic struct {
	TopicId int64  `json:"topic_id"`
	Group   Group  `json:"group"`
	Type    string `json:"type"`
	Talk    struct {
		Owner   User    `json:"owner"`
		Text    string  `json:"text"`
		Article Article `json:"article"`
	} `json:"talk"`
	LatestLikes []struct {
		CreateTime string `json:"create_time"`
		Owner      User   `json:"owner"`
	} `json:"latest_likes"`
	ShowComments []struct {
		CommentId       int64  `json:"comment_id"`
		CreateTime      string `json:"create_time"`
		Owner           User   `json:"owner"`
		Text            string `json:"text"`
		LikesCount      int    `json:"likes_count"`
		RewardsCount    int    `json:"rewards_count"`
		Sticky          bool   `json:"sticky"`
		RepliesCount    int    `json:"replies_count,omitempty"`
		ParentCommentId int64  `json:"parent_comment_id,omitempty"`
		Repliee         User   `json:"repliee,omitempty"`
	} `json:"show_comments"`
	LikesCount    int    `json:"likes_count"`
	RewardsCount  int    `json:"rewards_count"`
	CommentsCount int    `json:"comments_count"`
	ReadingCount  int    `json:"reading_count"`
	ReadersCount  int    `json:"readers_count"`
	Digested      bool   `json:"digested"`
	Sticky        bool   `json:"sticky"`
	CreateTime    string `json:"create_time"`
	UserSpecific  struct {
		Liked      bool `json:"liked"`
		Subscribed bool `json:"subscribed"`
	} `json:"user_specific"`
	Title   string `json:"title"`
	Columns []struct {
		ColumnId int64  `json:"column_id"`
		Name     string `json:"name"`
	} `json:"columns"`
}

type TopicInfo struct {
	Topic Topic  `json:"topic"`
	Type  string `json:"type"`
}
type Comment struct {
	CommentId       int64   `json:"comment_id"`
	CreateTime      string  `json:"create_time"`
	Owner           User    `json:"owner"`
	Text            string  `json:"text"`
	LikesCount      int     `json:"likes_count"`
	RewardsCount    int     `json:"rewards_count"`
	Sticky          bool    `json:"sticky"`
	Images          []Image `json:"images,omitempty"`
	RepliesCount    int     `json:"replies_count,omitempty"`
	RepliedComments []struct {
		CommentId       int64  `json:"comment_id"`
		ParentCommentId int64  `json:"parent_comment_id"`
		CreateTime      string `json:"create_time"`
		Owner           User   `json:"owner"`
		Text            string `json:"text"`
		LikesCount      int    `json:"likes_count"`
		RewardsCount    int    `json:"rewards_count"`
		Sticky          bool   `json:"sticky"`
		Repliee         User   `json:"repliee"`
	} `json:"replied_comments,omitempty"`
}

type Comments struct {
	Comments []Comment `json:"comments"`
}
