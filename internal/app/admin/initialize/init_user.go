package initialize

import (
	"lime/internal/app/admin/model"
	"lime/internal/app/admin/service"
	"lime/internal/common"
	"lime/internal/global"
)

// 初始化管理员用户数据
func InitAdminUser() {
	// 检查是否已经存在管理员用户
	adminUser, err := service.NewUser().GetUserByName("admin")
	if err != nil {
		return
	}

	if adminUser.ID != 0 {
		return
	}

	// 创建管理员用户
	password, err := common.MakePassword(global.DEFAULT_PWD)
	if err != nil {
		global.Logger.Sugar().Errorf("生成密码哈希失败: %s", err.Error())
		return
	}

	user := model.User{
		Username:       "admin",
		Password:       password,
		NickName:       "管理员",
		Sex:            common.EST_MALE,
		Mobile:         "19999999999",
		Email:          "admin@lime.xyz",
		Avatar:         "https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png",
		IsCanNotRemove: 1,
	}

	err = service.NewUser().AddUser(&user)
	if err != nil {
		global.Logger.Sugar().Errorf("创建管理员用户失败: %s", err.Error())
		return
	}
}
