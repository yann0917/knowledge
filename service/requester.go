package service

import "io"

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

// GetTopics
// scope= all 最新, digests 精华, by_owner 只看星主, questions 问答,with_files 文件
func (s *Service) GetTopics(groupID, scope, count, endTime string) (list Topics, err error) {
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
