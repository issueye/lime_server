package config

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Data struct {
	Val any `json:"value"`
}

func NewData(value any) *Data {
	return &Data{Val: value}
}

func (data Data) Value() (driver.Value, error) {
	return json.Marshal(data)
}

func (j *Data) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan JSONData: %v", value)
	}
	return json.Unmarshal(bytes, j)
}

func ToData(value string) *Data {
	data := &Data{}
	err := json.Unmarshal([]byte(value), data)
	if err != nil {
		// logger.Error(err)
		return data
	}

	return data
}

func (data *Data) ToJson() string {
	name, err := json.Marshal(data)
	if err != nil {
		return ""
	}

	return string(name)
}

func (data *Data) String() string {
	return fmt.Sprintf("%v", data.Val)
}

// Param
// 参数
type Param struct {
	ID    int64  `gorm:"column:id;primaryKey;autoIncrement:false;type:int" json:"id" form:"id"` // 编码
	Group string `gorm:"column:group;size:50;" json:"group" form:"group"`                       // 分组
	Name  string `gorm:"column:name;size:255;" json:"name" form:"name"`                         // 参数名称
	Value *Data  `gorm:"column:value;size:255;" json:"value" form:"value"`                      // 参数值
	Mark  string `gorm:"column:mark;size:255;" json:"mark" form:"mark"`                         // 备注
}

func (Param) TableName() string {
	return "param_info"
}
