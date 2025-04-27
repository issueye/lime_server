package common

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type UintArr []uint

// 实现 go.sql 接口
func (u UintArr) Value() (driver.Value, error) {
	if len(u) == 0 {
		return []byte{}, nil
	}

	return json.Marshal(u)
}

// 实现 go.sql 接口
func (u *UintArr) Scan(v interface{}) error {
	if v == nil {
		*u = []uint{}
		return nil
	}

	// 这里的 v 是 []byte []uint8 或 string 类型
	switch t := v.(type) {
	case []byte:
		{
			if len(t) == 0 {
				*u = []uint{}
				return nil
			}
			return json.Unmarshal(t, u)
		}
	case string:
		{
			if t == "" {
				*u = []uint{}
				return nil
			}
			return json.Unmarshal([]byte(t), u)
		}
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}
}
