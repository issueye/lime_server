package common

import (
	"database/sql/driver"
	"encoding/json"
)

type UintArr []uint

// 实现 go.sql 接口
func (u UintArr) Value() (driver.Value, error) {
	if len(u) == 0 {
		return nil, nil
	}

	return json.Marshal(u)
}

// 实现 go.sql 接口
func (u *UintArr) Scan(v interface{}) error {
	if v == nil {
		*u = []uint{}
		return nil
	}

	switch v := v.(type) {
	case []byte:
		return json.Unmarshal(v, u)
	case string:
		return json.Unmarshal([]byte(v), u)
	default:
		return nil
	}
}
