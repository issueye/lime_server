package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strings"
)

type Arr []string

// 存入数据库
func (arr Arr) Value() (driver.Value, error) {
	if len(arr) > 0 {
		str := arr[0]
		for _, v := range arr[1:] {
			str += "," + v
		}
		return str, nil
	} else {
		return "", nil
	}
}

// 从数据库取数据
func (arr *Arr) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("不匹配的数据类型")
	}
	*arr = strings.Split(string(str), ",")
	return nil
}

type KV struct {
	Key   string `json:"key"`   // 环境变量键
	Value string `json:"value"` // 环境变量值
}

type KVS []KV

func (h KVS) Value() (driver.Value, error) {
	if len(h) == 0 {
		return []byte{}, nil
	}
	return json.Marshal(h)
}

// Scan 从数据库中读取 JSON 数据并解析为 Headers
func (h *KVS) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var data []byte
	switch v := value.(type) {
	case string:
		data = []byte(v)
	case []byte:
		data = v
	default:
		return errors.New("不支持的数据类型")
	}

	if len(data) == 0 || string(data) == "{}" {
		return nil
	}

	return json.Unmarshal(data, h)
}
