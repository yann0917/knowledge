package service

import "io"

func (s *Service) GetSettings() (io.ReadCloser, error) {
	resp, err := s.client.R().
		Get("/settings")
	return handleHTTPResponse(resp, err)
}

func (s *Service) GetUserSelf() (io.ReadCloser, error) {
	resp, err := s.client.R().
		Get("/users/self")
	return handleHTTPResponse(resp, err)
}

func (s *Service) GetGroups() (io.ReadCloser, error) {
	resp, err := s.client.R().
		Get("/groups")
	return handleHTTPResponse(resp, err)
}

func (s *Service) GetMenus(groupID string) (io.ReadCloser, error) {
	resp, err := s.client.R().
		SetPathParams(map[string]string{
			"group_id": groupID,
		}).
		Get("/groups/{group_id}/menus")
	return handleHTTPResponse(resp, err)
}

// GetTopics
// scope= all 最新, digests 精华, by_owner 只看星主, questions 问答,with_files 文件
func (s *Service) GetTopics(groupID, scope, count, endTime string) (io.ReadCloser, error) {
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
	return handleHTTPResponse(resp, err)
}
