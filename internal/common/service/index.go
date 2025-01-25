package service

import (
	"errors"
	"lime/internal/common/model"
	"lime/internal/global"

	"gorm.io/gorm"
)

type Callback[T any] func(T, *gorm.DB) *gorm.DB

type BaseService[T any] struct {
	IsTX bool     // 是否开启事务
	TX   *gorm.DB // 事务对象
	DB   *gorm.DB // 数据库对象
}

func (b *BaseService[T]) GetDB() *gorm.DB {
	if b.TX != nil {
		return b.TX
	}

	return b.DB
}

func NewSrv[T any](base BaseService[T], args ...any) BaseService[T] {

	var (
		gdb  *gorm.DB
		isTx bool
	)

	if len(args) > 0 {
		switch args[0].(type) {
		case *gorm.DB:
			gdb = args[0].(*gorm.DB)
		}
	}

	if len(args) > 1 {
		switch args[1].(type) {
		case bool:
			isTx = args[1].(bool)
		}
	}

	if gdb != nil {
		if isTx {
			base.IsTX = true
			base.TX = gdb
		} else {
			base.DB = gdb
		}
	} else {
		base.DB = global.DB
	}

	return base
}

func (b *BaseService[T]) Begin() {
	if !b.IsTX {
		b.IsTX = true
		b.TX = b.DB.Begin()
	}
}

func (b *BaseService[T]) Rollback() error {
	if !b.IsTX {
		return errors.New("当前不是事务")
	}

	return b.TX.Rollback().Error
}

func (b *BaseService[T]) Commit() error {
	if !b.IsTX {
		return errors.New("当前不是事务")
	}

	return b.TX.Commit().Error
}

// 根据ID查询数据
func (b *BaseService[T]) GetById(id uint) (*T, error) {
	data := new(T)
	err := b.GetDB().Model(data).Where("id = ?", id).Find(data).Error
	return data, err
}

// 根据字段查询
func (b *BaseService[T]) GetByField(field string, value any) (*T, error) {
	data := new(T)
	err := b.GetDB().Model(data).Where(field+" = ?", value).Find(data).Error
	return data, err
}

// 根据字段查询
func (b *BaseService[T]) GetByMap(conditions map[string]any) (*T, error) {
	data := new(T)
	qry := b.GetDB().Model(data)
	for k, v := range conditions {
		qry = qry.Where(k+" =?", v)
	}
	err := qry.Find(data).Error
	return data, err
}

// 根据字段查询
func (b *BaseService[T]) GetDatasByField(field string, value any) ([]*T, error) {
	data := make([]*T, 0)
	err := b.GetDB().Model(new(T)).Where(field+" = ?", value).Find(&data).Error
	return data, err
}

// 根据字段查询
func (b *BaseService[T]) GetDatasByMap(conditions map[string]any) ([]*T, error) {
	data := make([]*T, 0)

	qry := b.GetDB().Model(new(T))
	for k, v := range conditions {
		qry = qry.Where(k+" =?", v)
	}

	err := qry.Find(&data).Error
	return data, err
}

// 根据ID更新数据(map)
func (b *BaseService[T]) UpdateByMap(id uint, data map[string]any) error {
	return b.GetDB().Model(new(T)).Where("id = ?", id).Updates(data).Error
}

func (b *BaseService[T]) UpdateByField(field string, value any, data map[string]any) error {
	return b.GetDB().Model(new(T)).Where(field+" = ?", value).Updates(data).Error
}

func (b *BaseService[T]) UpdatedatasByMap(conditions map[string]any, data map[string]any) error {
	srv := b.GetDB().Model(new(T))
	for k, v := range conditions {
		srv = srv.Where(k+" =?", v)
	}

	return srv.Updates(data).Error
}

// 更新ID更新数据(结构体)
func (b *BaseService[T]) Update(id uint, data *T) error {
	return b.GetDB().Model(new(T)).Where("id = ?", id).Updates(data).Error
}

// 创建数据
func (b *BaseService[T]) Create(data *T) error {
	return b.GetDB().Create(data).Error
}

// 批量创建数据
func (b *BaseService[T]) CreateBatch(data []T) error {
	return b.GetDB().Create(&data).Error
}

// 删除数据
func (b *BaseService[T]) Delete(id uint) error {
	return b.GetDB().Model(new(T)).Where("id = ?", id).Delete(new(T)).Error
}

// 保存数据
func (b *BaseService[T]) SaveByCondition(condition map[string]any, data T) error {
	srv := b.GetDB().Model(new(T))
	for k, v := range condition {
		srv = srv.Where(k+" =?", v)
	}

	return srv.Save(&data).Error
}

func (b *BaseService[T]) DeleteByFields(condition map[string]any) error {
	del := b.GetDB().Model(new(T))

	for k, v := range condition {
		del = del.Where(k+" = ?", v)
	}

	return del.Delete(new(T)).Error
}

func (b *BaseService[T]) GetAllCount() (int64, error) {
	var count int64
	err := b.GetDB().Model(new(T)).Count(&count).Error
	return count, err
}

func (b *BaseService[T]) GetCountByFields(condition map[string]any) (int64, error) {
	var count int64
	qry := b.GetDB().Model(new(T))

	for k, v := range condition {
		qry = qry.Where(k+" = ?", v)
	}

	err := qry.Count(&count).Error
	return count, err
}

// 获取列表
func (b *BaseService[T]) GetAll() ([]*T, error) {
	var data []*T
	err := b.GetDB().Model(new(T)).Find(&data).Error
	return data, err
}

func GetList[T any, F any](condition *model.PageQuery[F], callback Callback[F]) (*model.ResPage[T], error) {
	var data []*T
	qry := global.DB.Model(new(T))
	qry = callback(condition.Condition, qry)

	count := int64(0)
	err := qry.Count(&count).Error
	if err != nil {
		return nil, err
	}

	condition.Total = int(count)

	// 进行分页查询
	if condition.PageNum == 0 || condition.PageSize == 0 {
		err = qry.Find(&data).Error
	} else {
		err = qry.
			Limit(condition.PageSize).
			Offset((condition.PageNum - 1) * condition.PageSize).
			Find(&data).Error
	}

	if err != nil {
		return nil, err
	}

	rtnData := model.NewResPage(condition.PageNum, condition.PageSize, condition.Total, data)
	return rtnData, err
}

func GetDatas[T any, F any](condition F, callback Callback[F]) ([]*T, error) {
	var data []*T
	qry := global.DB.Model(new(T))
	qry = callback(condition, qry)

	count := int64(0)
	err := qry.Count(&count).Error
	if err != nil {
		return nil, err
	}

	err = qry.Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, err
}
