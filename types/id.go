package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
)

// ID 是一个基于 int64 的自定义类型，
// JSON 序列化时为 string，数据库中存为 int64。
type ID int64

func IDFrom(id int64) ID {
	return (ID)(id)
}

func (s ID) ToInt64() int64 {
	return (int64)(s)
}

func (s ID) ToString() string {
	return strconv.FormatInt(int64(s), 10)
}

func (s ID) Equals(id ID) bool {
	return s.ToInt64() == id.ToInt64()
}

func (s ID) EqualsInt64(id int64) bool {
	return s.ToInt64() == id
}

func (s ID) Compare(id ID) int {
	return int(s.ToInt64() - id.ToInt64())
}

func IDCompare(id1 ID, id2 ID) int {
	return id1.Compare(id2)
}

// MarshalJSON 实现 json.Marshaler 接口
// 在序列化时转为字符串
func (id ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(strconv.FormatInt(int64(id), 10))
}

// UnmarshalJSON 实现 json.Unmarshaler 接口
// 支持 string 或 number 两种 JSON 表示
func (id *ID) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err == nil {
		v, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid ID string: %v", err)
		}
		*id = ID(v)
		return nil
	}

	// 如果不是字符串，尝试直接解析为数字
	var v int64
	if err := json.Unmarshal(b, &v); err != nil {
		return fmt.Errorf("invalid ID value: %v", err)
	}
	*id = ID(v)
	return nil
}

// Value 实现 driver.Valuer 接口（用于写入数据库）
func (id ID) Value() (driver.Value, error) {
	return int64(id), nil
}

// Scan 实现 sql.Scanner 接口（用于读取数据库）
func (id *ID) Scan(value any) error {
	switch v := value.(type) {
	case int64:
		*id = ID(v)
	case int:
		*id = ID(v)
	case []byte:
		i, err := strconv.ParseInt(string(v), 10, 64)
		if err != nil {
			return err
		}
		*id = ID(i)
	case string:
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return err
		}
		*id = ID(i)
	default:
		return fmt.Errorf("unsupported Scan type for ID: %T", value)
	}
	return nil
}

// String 实现 Stringer 接口
func (id ID) String() string {
	return strconv.FormatInt(int64(id), 10)
}
