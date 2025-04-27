package model

import (
	"lime/internal/common"
	"lime/internal/common/model"
)

type User struct {
	model.BaseModel
	Username       string             `gorm:"column:username;unique;size:50;not null;comment:用户名，用于登录和识别用户;" json:"username"`
	Password       string             `gorm:"column:password;size:255;not null;comment:用户密码，以加密形式存储;" json:"password"`
	NickName       string             `gorm:"column:nick_name;size:50;not null;comment:用户昵称，用于显示在页面上，可修改;" json:"nick_name"`
	Avatar         string             `gorm:"column:avatar;size:255;not null;comment:用户头像，用于显示在页面上，可修改;" json:"avatar"`
	Sex            common.EnumSexType `gorm:"column:sex;not null;comment:用户性别，0 未知 1 男 2 女;" json:"sex"`
	Mobile         string             `gorm:"column:mobile;size:20;not null;comment:用户手机号，用于接收通知和联系用户;" json:"mobile"`
	Email          string             `gorm:"column:email;size:50;not null;comment:用户邮箱，用于接收通知和联系用户;" json:"email"`
	Remark         string             `gorm:"column:remark;size:255;not null;comment:用户备注，用于描述用户;" json:"remark"`
	IsCanNotRemove uint               `gorm:"column:is_can_not_remove;comment:是否可删除;" json:"is_can_not_remove"` // 是否可删除 0 否 1 是
	// 一个操作员只能有一个角色
	UserRole *UserRole `gorm:"foreignKey:user_id"`
}

func (u User) TableName() string { return "sys_user" }
