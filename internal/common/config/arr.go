package config

import (
	"database/sql/driver"
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

func (arr *Arr) UnmarshalJSON(data []byte) error {
	d := string(data)

	if d == "null" {
		*arr = []string{}
		return nil
	}

	// 处理 ["983193","983173"] 这种情况
	if strings.HasPrefix(d, "[") || strings.HasPrefix(d, "]") {
		d = strings.Trim(d, "[")
		d = strings.Trim(d, "]")
	}

	if d != "" {
		arrData := strings.Split(d, ",")
		for _, data := range arrData {
			data = strings.Trim(data, "\"")
			*arr = append(*arr, data)
		}
	} else {
		*arr = []string{}
	}

	return nil
}
