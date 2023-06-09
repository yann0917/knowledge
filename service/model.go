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
	GroupId       int64  `json:"group_id"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	BackgroundURL string `json:"background_url"`
}

type GroupInfoResp struct {
	Group GroupInfo `json:"group"`
}

type GroupInfo struct {
	GroupId                          int64  `json:"group_id"`
	Number                           int    `json:"number"`
	Name                             string `json:"name"`
	Description                      string `json:"description"`
	CreateTime                       string `json:"create_time"`
	UpdateTime                       string `json:"update_time"`
	PrivilegeUserLastTopicCreateTime string `json:"privilege_user_last_topic_create_time"`
	LatestTopicCreateTime            string `json:"latest_topic_create_time"`
	AliveTime                        string `json:"alive_time"`
	BackgroundUrl                    string `json:"background_url"`
	Type                             string `json:"type"`
	RiskLevel                        string `json:"risk_level"`
	Category                         struct {
		CategoryId int    `json:"category_id"`
		Title      string `json:"title"`
	} `json:"category"`
	Owner      User          `json:"owner"`
	AdminIds   []int64       `json:"admin_ids"`
	GuestIds   []int64       `json:"guest_ids"`
	PartnerIds []int64       `json:"partner_ids"`
	Promos     []interface{} `json:"promos"`
	Policies   struct {
		NeedExamine             bool   `json:"need_examine"`
		EnableScoreboard        bool   `json:"enable_scoreboard"`
		FreeQuestionsLimitCount int    `json:"free_questions_limit_count"`
		EnableMemberNumber      bool   `json:"enable_member_number"`
		MembersVisibility       string `json:"members_visibility"`
		AllowSharing            bool   `json:"allow_sharing"`
		AllowPrivateChat        bool   `json:"allow_private_chat"`
		AllowSearch             bool   `json:"allow_search"`
		AllowJoin               bool   `json:"allow_join"`
		AllowAnonymousQuestion  bool   `json:"allow_anonymous_question"`
		SilenceNewMember        bool   `json:"silence_new_member"`
		EnableWatermark         bool   `json:"enable_watermark"`
		ParseBookTitle          bool   `json:"parse_book_title"`
		AllowCopy               bool   `json:"allow_copy"`
		AllowDownload           bool   `json:"allow_download"`
		EnableIap               bool   `json:"enable_iap"`
	} `json:"policies"`
	Privileges struct {
		AccessGroupData     string `json:"access_group_data"`
		AccessIncomesData   string `json:"access_incomes_data"`
		AccessWeeklyReports string `json:"access_weekly_reports"`
		CreateTopic         string `json:"create_topic"`
		CreateComment       string `json:"create_comment"`
	} `json:"privileges"`
	Statistics struct {
		Topics struct {
			TopicsCount  int `json:"topics_count"`
			AnswersCount int `json:"answers_count"`
			DigestsCount int `json:"digests_count"`
		} `json:"topics"`
		Files struct {
			Count int `json:"count"`
		} `json:"files"`
		Members struct {
			Count int `json:"count"`
		} `json:"members"`
	} `json:"statistics"`
	UserSpecific struct {
		JoinTime        string `json:"join_time"`
		RewardedOwner   bool   `json:"rewarded_owner"`
		EnableFootprint bool   `json:"enable_footprint"`
	} `json:"user_specific"`
}

type Groups struct {
	Groups []Group `json:"groups"`
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

type UserInfo struct {
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
	ArticleID        string `json:"article_id"`
	ArticleURL       string `json:"article_url"`
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

type Question struct {
	Owner  User    `json:"owner"`  // 提问者
	Text   string  `json:"text"`   // 问题内容
	Images []Image `json:"images"` // 问题图片
}

type Answer struct {
	Owner  User    `json:"owner"`  // 回答者
	Text   string  `json:"text"`   // 回答内容
	Images []Image `json:"images"` // 回答图片
}

type ShowComment struct {
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
}

type Topic struct {
	Question Question `json:"question"` // 提问
	Answer   Answer   `json:"answer"`   // 回答
	TopicId  int64    `json:"topic_id"`
	Group    Group    `json:"group"`
	Type     string   `json:"type"`
	Talk     struct {
		Owner   User    `json:"owner"`
		Text    string  `json:"text"`
		Article Article `json:"article,omitempty"`
		Images  []Image `json:"images,omitempty"`
		Files   []File  `json:"files,omitempty"`
	} `json:"talk"`
	// LatestLikes []struct {
	// 	CreateTime string `json:"create_time"`
	// 	Owner      User   `json:"owner"`
	// } `json:"latest_likes"`
	// ShowComments  []ShowComment `json:"show_comments"`
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

type Topics struct {
	Topics []Topic `json:"topics"`
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

type ColumnTopic struct {
	TopicId              int64  `json:"topic_id"`
	Title                string `json:"title"`
	Text                 string `json:"text"`
	CreateTime           string `json:"create_time"`
	AttachedToColumnTime string `json:"attached_to_column_time"`
}

type ColumnTopics struct {
	Topics []ColumnTopic `json:"topics,omitempty"`
}
