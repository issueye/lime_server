package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Script struct {
	Name    string `json:"name"`    // 脚本名称
	Content string `json:"content"` // 脚本内容
}

type Scripts []Script

func (h Scripts) Value() (driver.Value, error) {
	if len(h) == 0 {
		return []byte{}, nil
	}
	return json.Marshal(h)
}

// Scan 从数据库中读取 JSON 数据并解析为 Headers
func (h *Scripts) Scan(value interface{}) error {
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

type OS_TYPE int

const (
	OS_WINDOWS OS_TYPE = iota
	OS_LINUX
	OS_DARWIN
)

func (osType OS_TYPE) String() string {
	switch osType {
	case OS_LINUX:
		return "linux"
	case OS_DARWIN:
		return "darwin"
	default:
		return "windows"
	}
}

type ARCH_TYPE int

const (
	ARCH_AMD64 ARCH_TYPE = iota
	ARCH_ARM64
)

func (archType ARCH_TYPE) String() string {
	switch archType {
	case ARCH_ARM64:
		return "arm64"
	default:
		return "amd64"
	}
}
