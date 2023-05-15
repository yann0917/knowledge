package base

import (
	"io"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
)

const UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36"
const format = "2006-01-02T15:04:05.000+0800"
const TimeZone = "Asia/Shanghai"

func GetParseParam(param string) string {
	return url.PathEscape(param)
}

func TimeToStr(t time.Time) string {
	return t.Format(format)
}

func StrToTime(s string) time.Time {
	t, _ := time.Parse(format, s)
	return t
}

func StrToTimeUtf8(s string) time.Time {
	var cst0, _ = time.LoadLocation(TimeZone)
	t, _ := time.ParseInLocation(format, s, cst0)
	return t
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func GetParentDirectory(directory string) string {
	return substr(directory, 0, strings.LastIndex(directory, "/"))
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
		// log.Info(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

// UnmarshalReader 将 r 中的 json 格式的数据, 解析到 v
func UnmarshalReader(r io.Reader, v interface{}) error {
	d := jsoniter.NewDecoder(r)
	return d.Decode(v)
}

// UnmarshalJSON 将 data 中的 json 格式的数据, 解析到 v
func UnmarshalJSON(data []byte, v interface{}) error {
	return jsoniter.Unmarshal(data, v)
}
