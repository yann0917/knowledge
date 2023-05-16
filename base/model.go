package base

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"

type Model struct {
	// cache *cache.Cache
}

// JSONTime format json time field by myself
type JSONTime struct {
	time.Time
}

func (t *JSONTime) UnmarshalJSON(data []byte) (err error) {
	if string(data) == "null" || string(data) == `""` {
		return
	}
	if len(data) == 2 {
		*t = JSONTime{Time: time.Time{}}
		return
	}
	loc, _ := time.LoadLocation("Asia/Shanghai")
	now, err := time.ParseInLocation(`"`+TimeFormat+`"`, string(data), loc)
	*t = JSONTime{Time: now}
	return
}

// MarshalJSON on JSONTime format Time field with Y-m-d H:i:s
func (t JSONTime) MarshalJSON() ([]byte, error) {
	if t.Time.IsZero() {
		return []byte("null"), nil
	}
	formatted := fmt.Sprintf("\"%s\"", t.Format(TimeFormat))
	return []byte(formatted), nil
}

// Value insert timestamp into mysql need this function.
func (t JSONTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan value of time.Time
func (t *JSONTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = JSONTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// ValidateJSONTimeType gin 验证器 binding:"required"
func ValidateJSONTimeType(field reflect.Value) interface{} {
	if field.Type() == reflect.TypeOf(JSONTime{}) {
		timeStr := field.Interface().(JSONTime).String()
		// 这里返回 Nil 则会被 validator 判定为空值，而无法通过 `binding:"required"` 规则
		if field.Interface().(JSONTime).Time.IsZero() {
			return nil
		}
		return timeStr
	}
	return nil
}
