package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type KeyValue struct {
	Key    string `json:"key"`
	Data   string `json:"data"`
	Remark string `json:"remark"`
}

type KVs []KeyValue

func (data *KVs) Scan(value interface{}) error {
	// 如果数据为空
	if value == nil {
		return nil
	}

	return json.Unmarshal(value.([]byte), data)
}

func (data *KVs) Value() (driver.Value, error) {
	// 如果数据为空
	if data == nil {
		return nil, nil
	}

	return json.Marshal(data)
}

type BaseModel struct {
	ID        uint      `gorm:"column:id;primaryKey;not null" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;" json:"updated_at"`
}

type BasePage struct {
	PageNum  int `json:"pageNum" form:"pageNum"`
	PageSize int `json:"pageSize" form:"pageSize"`
	Total    int `json:"total" form:"total"`
}

type ResPage[T any] struct {
	BasePage
	List []*T `json:"list"`
}

func NewResPage[T any](pageNum, pageSize int, total int, list []*T) *ResPage[T] {
	return &ResPage[T]{
		BasePage: BasePage{PageNum: pageNum, PageSize: pageSize, Total: total},
		List:     list,
	}
}

type PageQuery[T any] struct {
	BasePage
	Condition T `json:"condition"`
}

func NewPageQuery[T any](pageNum, pageSize int, condition T) *PageQuery[T] {
	return &PageQuery[T]{
		BasePage:  BasePage{PageNum: pageNum, PageSize: pageSize},
		Condition: condition,
	}
}
