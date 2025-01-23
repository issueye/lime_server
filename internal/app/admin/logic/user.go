package logic

import (
	"errors"
	"lime/internal/app/admin/model"
	"lime/internal/app/admin/requests"
	"lime/internal/app/admin/service"
	commonModel "lime/internal/common/model"
	"lime/internal/global"
)

func ListUser(condition *commonModel.PageQuery[*requests.QueryUser]) (*commonModel.ResPage[model.User], error) {
	return service.NewUser().ListUser(condition)
}

func UpdateUser(u *model.User) error {
	data := make(map[string]any)
	data["username"] = u.Username
	data["nick_name"] = u.NickName
	data["avatar"] = u.Avatar

	return service.NewUser().UpdateByMap(u.ID, data)
}

func UpdatePassword(user model.User, u *requests.UpdatePassword) error {
	// 检查旧密码是否正确
	userInfo, err := service.NewUser().GetById(user.ID)
	if err != nil {
		return err
	}

	// 检查两次密码是否一致
	if u.Password != u.Repassword {
		return errors.New("两次密码不一致")
	}

	// 加密密码
	pwd, err := MakePassword(u.Oldpassword)
	if err != nil {
		return err
	}

	if userInfo.Password != pwd {
		return errors.New("旧密码错误")
	}

	// 加密密码
	pwd, err = MakePassword(u.Password)
	if err != nil {
		return err
	}

	// 更新密码
	return service.NewUser().UpdateByMap(user.ID, map[string]any{"password": pwd})
}

func UpdateUserInfo(u *requests.UpdateUser) error {
	data := make(map[string]any)
	data["username"] = u.Username
	data["nick_name"] = u.NickName
	data["avatar"] = u.Avatar

	return service.NewUser().UpdateByMap(uint(u.Id), data)
}

func DeleteUser(id uint) error {
	return service.NewUser().Delete(id)
}

func GetUserById(id uint) (*model.User, error) {
	return service.NewUser().GetById(id)
}

func CreateUser(u *requests.CreateUser) error {
	pwd, err := MakePassword(global.DEFAULT_PWD)
	if err != nil {
		return err
	}

	info := &model.User{
		Username: u.Username,
		NickName: u.NickName,
		Avatar:   u.Avatar,
		Password: pwd,
		UserRole: &model.UserRole{
			RoleCode: u.RoleCode,
		},
	}

	return service.NewUser().AddUser(info)
}

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
	password, err := MakePassword(global.DEFAULT_PWD)
	if err != nil {
		global.Logger.Sugar().Errorf("生成密码哈希失败: %s", err.Error())
		return
	}

	user := model.User{
		Username: "admin",
		Password: password,
		NickName: "管理员",
		Avatar:   "https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png",
	}

	err = service.NewUser().AddUser(&user)
	if err != nil {
		global.Logger.Sugar().Errorf("创建管理员用户失败: %s", err.Error())
		return
	}
}
