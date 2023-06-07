package base

import (
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

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

func StrToTimeUtc8(s string) time.Time {
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

// String2Int 字符串 转 int
func String2Int(str string) (int, error) {
	return strconv.Atoi(str)
}

// String2Int64 字符串 转 int64
func String2Int64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

// String2Uint string 转 uint
func String2Uint(str string) uint {
	u64, _ := strconv.ParseUint(str, 10, 0)
	return uint(u64)
}

// Int2String int 转 string
func Int2String(intval int) string {
	return strconv.Itoa(intval)
}

// Int642String int64 转 string
func Int642String(intval int64) string {
	return strconv.FormatInt(intval, 10)
}

func Uint642String(intval uint64) string {
	return strconv.FormatUint(intval, 10)
}

// Byte2String 字节数组转 string
func Byte2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// Float642String float64 转 string
func Float642String(v float64, p int) string {
	return strconv.FormatFloat(v, 'f', p, 64)
}

// String2Float64 string 转 float64
func String2Float64(str string) (float64, error) {
	return strconv.ParseFloat(str, 64)
}

func Bool2Int(b bool) int {
	if b {
		return 1
	}
	return 0
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

// FileName filter invalid string
func FileName(name, ext string) string {
	rep := strings.NewReplacer("\n", " ", "/", " ", "|", "-", ": ", "：", ":", "：", "'", "’", "\t", " ")
	name = rep.Replace(name)

	if runtime.GOOS == "windows" {
		rep = strings.NewReplacer("\"", " ", "?", " ", "*", " ", "\\", " ", "<", " ", ">", " ", ":", " ", "：", " ")
		name = rep.Replace(name)
	}

	name = strings.TrimSpace(name)

	limitedName := LimitLength(name, 80)
	if ext != "" {
		return fmt.Sprintf("%s.%s", limitedName, ext)
	}
	return limitedName
}

// LimitLength cut string
func LimitLength(s string, length int) string {
	ellipses := "..."

	str := []rune(s)
	if len(str) > length {
		s = string(str[:length-len(ellipses)]) + ellipses
	}

	return s
}

// FilePath gen valid file path
func FilePath(name, ext string, escape bool) (string, error) {
	var outputPath string

	fileName := name
	if escape {
		fileName = FileName(name, ext)
	} else {
		if ext != "" {
			fileName = fmt.Sprintf("%s.%s", name, ext)
		}
	}
	outputPath = filepath.Join(fileName)
	return outputPath, nil
}

// Mkdir mkdir path
func Mkdir(elem ...string) (string, error) {
	path := filepath.Join(elem...)

	err := os.MkdirAll(path, os.ModePerm)

	return path, err
}

// FileSize return the file size of the specified path file
func FileSize(filePath string) (int, bool, error) {
	file, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return 0, false, nil
		}
		return 0, false, err
	}
	return int(file.Size()), true, nil
}

func GetMdHeader(level int) string {
	heads := map[int]string{
		1: "# ",
		2: "## ",
		3: "### ",
		4: "#### ",
		5: "##### ",
		6: "###### ",
	}
	if s, ok := heads[level]; ok {
		return s
	}
	return ""
}

func SaveFile(name, content string) (err error) {
	path, err := Mkdir(GetCurrentDirectory())
	if err != nil {
		return
	}
	fileName := filepath.Join(path, name)
	// _, exist, err := FileSize(fileName)
	// if exist {
	// 	fmt.Printf("\033[33;1m%s\033[0m\n", "已存在")
	// }
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return
	}
	_, err = f.WriteString(content)
	if err != nil {
		return
	}
	if err = f.Close(); err != nil {
		return
	}
	return
}

func CheckFileExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func WriteFileWithTrunc(filename, content string) (err error) {

	var f *os.File
	if CheckFileExist(filename) {
		f, err = os.OpenFile(filename, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0666)

		if err != nil {
			return
		}
	} else {
		f, err = os.Create(filename)
		if err != nil {
			return
		}
	}
	defer f.Close()
	_, err = io.WriteString(f, content)
	if err != nil {
		return
	}

	err = f.Sync()
	return

}
