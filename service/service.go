package service

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/mitchellh/mapstructure"
	"github.com/yann0917/knowledge/base"
)

var (
	BaseURL   = "https://api.zsxq.com/v2"
	UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36"
)

// Response response
type Response struct {
	Succeeded bool   `json:"succeeded"`
	Code      int    `json:"code,omitempty"`
	Info      string `json:"info,omitempty"`
	RespData  respC  `json:"resp_data"`
}

// respC response content
type respC []byte

func (r *respC) UnmarshalJSON(data []byte) error {
	*r = data

	return nil
}

func (r respC) String() string {
	return string(r)
}

// Service dedao service
type Service struct {
	client *resty.Client
}

// CookieOptions dedao cookie options
type CookieOptions struct {
	Zsxqsessionid   string `json:"zsxqsessionid"`
	AbtestEnv       string `json:"abtest_env" mapstructure:"abtest_env"`
	ZsxqAccessToken string `json:"zsxq_access_token" mapstructure:"zsxq_access_token"`
	UabCollina      string `json:"_uab_collina" mapstructure:"_uab_collina"`
	Debug           bool   `json:"debug"`
}

// NewService new service
func NewService(co *CookieOptions) *Service {
	cookies := []*http.Cookie{}
	cookies = append(cookies, &http.Cookie{
		Name:   "zsxqsessionid",
		Value:  co.Zsxqsessionid,
		Domain: "api.zsxq.com",
	})
	cookies = append(cookies, &http.Cookie{
		Name:   "abtest_env",
		Value:  co.AbtestEnv,
		Domain: ".zsxq.com",
	})
	cookies = append(cookies, &http.Cookie{
		Name:   "zsxq_access_token",
		Value:  co.ZsxqAccessToken,
		Domain: ".zsxq.com",
	})
	cookies = append(cookies, &http.Cookie{
		Name:   "_uab_collina",
		Value:  co.UabCollina,
		Domain: "wx.zsxq.com",
	})
	client := resty.New()
	if co.Debug {
		client.SetDebug(true)
	}
	client.SetBaseURL(BaseURL).
		SetCookies(cookies).
		SetHeader("User-Agent", UserAgent).
		SetHeader("Accept", "application/json, text/plain, */*").
		SetHeader("X-Timestamp", strconv.FormatInt(time.Now().Unix(), 10)).
		SetHeader("X-Version", "2.37.0")

	return &Service{client: client}
}

func (r *Response) isSuccess() bool {
	return r.Succeeded == true
}

func handleHTTPResponse(resp *resty.Response, err error) (io.ReadCloser, error) {
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, errors.New("404 NotFound")
	}
	if resp.StatusCode() == http.StatusBadRequest {
		return nil, errors.New("400 BadRequest")
	}
	if resp.StatusCode() == http.StatusUnauthorized {
		return nil, errors.New("401 Unauthorized")
	}
	if resp.StatusCode() == 496 {
		return nil, errors.New("496 NoCertificate")
	}

	data := resp.Body()
	reader := bytes.NewReader(data)
	result := io.NopCloser(reader)
	return result, nil
}

func handleJSONParse(reader io.Reader, v interface{}) error {
	result := new(Response)

	err := base.UnmarshalReader(reader, &result)
	if err != nil {
		fmt.Printf("err1: %s \n", err.Error())
		return err
	}
	if !result.isSuccess() {
		err = errors.New("接口请求异常，请稍后重试。errMsg:" + result.Info)
		return err
	}
	err = base.UnmarshalJSON(result.RespData, v)
	if err != nil {
		fmt.Printf("err2: %s", err.Error())
		return err
	}

	return nil
}

// ParseCookies parse cookie string to cookie options
func ParseCookies(cookie string, v interface{}) (err error) {
	if cookie == "" {
		return errors.New("cookie is empty")
	}
	list := strings.Split(cookie, "; ")
	cookieM := make(map[string]string, len(list))
	for _, v := range list {
		item := strings.Split(v, "=")
		if len(item) > 1 {
			cookieM[item[0]] = item[1]
		}
	}
	err = mapstructure.Decode(cookieM, v)
	return
}
