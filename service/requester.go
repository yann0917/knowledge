package service

import (
	"io"

	"github.com/PuerkitoBio/goquery"
)

func (s *Service) GetSettings() (io.ReadCloser, error) {
	resp, err := s.client.R().
		Get("/settings")
	return handleHTTPResponse(resp, err)
}

func (s *Service) GetUserSelf() (user UserInfo, err error) {
	resp, err := s.client.R().
		Get("/users/self")
	if err != nil {
		return
	}
	body, err := handleHTTPResponse(resp, err)
	if err != nil {
		return
	}
	defer body.Close()
	if err = handleJSONParse(body, &user); err != nil {
		return
	}
	return
}

func (s *Service) GetGroups() (list Groups, err error) {
	resp, err := s.client.R().
		Get("/groups")
	if err != nil {
		return
	}
	body, err := handleHTTPResponse(resp, err)
	if err != nil {
		return
	}
	defer body.Close()
	if err = handleJSONParse(body, &list); err != nil {
		return
	}
	return
}

func (s *Service) GetMenus(groupID string) (list Menus, err error) {
	resp, err := s.client.R().
		SetPathParams(map[string]string{
			"group_id": groupID,
		}).
		Get("/groups/{group_id}/menus")
	if err != nil {
		return
	}
	body, err := handleHTTPResponse(resp, err)
	if err != nil {
		return
	}
	defer body.Close()
	if err = handleJSONParse(body, &list); err != nil {
		return
	}
	return
}

// GetGroupTopics
// scope= all 最新, digests 精华, by_owner 只看星主, questions 问答,with_files 文件
func (s *Service) GetGroupTopics(groupID, scope, count, endTime string) (list Topics, err error) {
	queryParam := map[string]string{
		"scope": scope,
		"count": count,
	}
	if endTime != "" {
		queryParam["end_time"] = endTime
	}
	resp, err := s.client.R().
		SetPathParams(map[string]string{
			"group_id": groupID,
		}).SetQueryParams(queryParam).
		Get("/groups/{group_id}/topics")
	if err != nil {
		return
	}
	body, err := handleHTTPResponse(resp, err)
	if err != nil {
		return
	}
	defer body.Close()
	if err = handleJSONParse(body, &list); err != nil {
		return
	}
	return
}

// GetColumns group columns 导读分类
func (s *Service) GetColumns(groupID string) (list Columns, err error) {
	resp, err := s.client.R().
		SetPathParams(map[string]string{
			"group_id": groupID,
		}).
		Get("/groups/{group_id}/columns")
	if err != nil {
		return
	}
	body, err := handleHTTPResponse(resp, err)
	if err != nil {
		return
	}
	defer body.Close()
	if err = handleJSONParse(body, &list); err != nil {
		return
	}
	return
}

// GetColumnTopics group columns topics 导读二级分类
func (s *Service) GetColumnTopics(groupID, columnID string) (list ColumnTopics, err error) {
	queryParam := map[string]string{
		"count":     "100",
		"sort":      "attached_to_column_time",
		"direction": "desc",
	}

	resp, err := s.client.R().
		SetPathParams(map[string]string{
			"group_id":  groupID,
			"column_id": columnID,
		}).SetQueryParams(queryParam).
		Get("/groups/{group_id}/columns/{column_id}/topics")
	if err != nil {
		return
	}
	body, err := handleHTTPResponse(resp, err)
	if err != nil {
		return
	}
	defer body.Close()
	if err = handleJSONParse(body, &list); err != nil {
		return
	}
	return
}

// GetTopicInfo group columns topics
func (s *Service) GetTopicInfo(topicID string) (info TopicInfo, err error) {
	resp, err := s.client.R().
		SetPathParams(map[string]string{
			"topic_id": topicID,
		}).
		Get("/topics/{topic_id}/info")
	if err != nil {
		return
	}
	body, err := handleHTTPResponse(resp, err)
	if err != nil {
		return
	}
	defer body.Close()
	if err = handleJSONParse(body, &info); err != nil {
		return
	}
	return
}

func (s *Service) GetArticle(url string) (info string, err error) {
	resp, err := s.client.SetBaseURL("").R().
		Get(url)
	if err != nil {
		return
	}
	body, err := handleHTTPResponse(resp, err)
	if err != nil {
		return
	}
	defer body.Close()
	// info =
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return
	}
	info, _ = doc.Find(".ql-editor").Html()
	return
}
