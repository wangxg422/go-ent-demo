package types

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type CustomTime time.Time

const timeFormat = "2006-01-02 15:04:05"

func NewCustomTimeNow() CustomTime {
	return CustomTime(time.Now())
}

// JSON 序列化
func (t CustomTime) MarshalJSON() ([]byte, error) {
	ts := time.Time(t).Format(timeFormat)
	return []byte(fmt.Sprintf("\"%s\"", ts)), nil
}

// JSON 反序列化
func (t *CustomTime) UnmarshalJSON(data []byte) error {
	str := string(data)
	str = str[1 : len(str)-1] // 去掉引号
	parsed, err := time.ParseInLocation(timeFormat, str, time.Local)
	if err != nil {
		return err
	}
	*t = CustomTime(parsed)
	return nil
}

// 数据库存取接口：driver.Valuer
func (t CustomTime) Value() (driver.Value, error) {
	return time.Time(t), nil
}

// 数据库存取接口：sql.Scanner
func (t *CustomTime) Scan(value any) error {
	if value == nil {
		*t = CustomTime(NewCustomTimeNow())
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		*t = CustomTime(v)
	case []byte:
		parsed, err := time.ParseInLocation(timeFormat, string(v), time.Local)
		if err != nil {
			return err
		}
		*t = CustomTime(parsed)
	default:
		return fmt.Errorf("unsupported type %T for CustomTime", value)
	}
	return nil
}

// 辅助函数：返回标准 time.Time
func (t CustomTime) Time() time.Time {
	return time.Time(t)
}
