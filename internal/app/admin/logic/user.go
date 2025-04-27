package logic

import (
	"errors"
	"lime/internal/app/admin/model"
	"lime/internal/app/admin/requests"
	"lime/internal/app/admin/service"
	"lime/internal/common"
	commonModel "lime/internal/common/model"
	"lime/internal/global"
)

type UserLogic struct{}

func NewUserLogic() *UserLogic {
	return &UserLogic{}
}

func (lc *UserLogic) ListUser(condition *commonModel.PageQuery[*requests.QueryUser]) (*commonModel.ResPage[model.User], error) {
	return service.NewUser().ListUser(condition)
}

func (lc *UserLogic) UpdateUser(u *model.User) error {
	data := make(map[string]any)
	data["username"] = u.Username
	data["nick_name"] = u.NickName
	data["sex"] = u.Sex
	data["avatar"] = u.Avatar
	data["mobile"] = u.Mobile
	data["email"] = u.Email
	data["remark"] = u.Remark

	return service.NewUser().UpdateByMap(u.ID, data)
}

func (lc *UserLogic) UpdatePassword(user model.User, u *requests.UpdatePassword) error {
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
	pwd, err := common.MakePassword(u.Oldpassword)
	if err != nil {
		return err
	}

	if userInfo.Password != pwd {
		return errors.New("旧密码错误")
	}

	// 加密密码
	pwd, err = common.MakePassword(u.Password)
	if err != nil {
		return err
	}

	// 更新密码
	return service.NewUser().UpdateByMap(user.ID, map[string]any{"password": pwd})
}

func (lc *UserLogic) UpdateUserInfo(u *requests.UpdateUser) error {
	data := make(map[string]any)
	data["username"] = u.Username
	data["nick_name"] = u.NickName
	data["sex"] = u.Sex
	data["mobile"] = u.Mobile
	data["email"] = u.Email
	data["remark"] = u.Remark
	data["avatar"] = u.Avatar

	return service.NewUser().UpdateByMap(uint(u.Id), data)
}

func (lc *UserLogic) DeleteUser(id uint) error {
	return service.NewUser().Delete(id)
}

func (lc *UserLogic) GetUserById(id uint) (*model.User, error) {
	return service.NewUser().GetById(id)
}

func (lc *UserLogic) ResetPwd(id uint) error {
	pwd, err := common.MakePassword(global.DEFAULT_PWD)
	if err != nil {
		return err
	}

	return service.NewUser().UpdateByMap(id, map[string]any{"password": pwd})
}

func (lc *UserLogic) CreateUser(u *requests.CreateUser) error {
	pwd, err := common.MakePassword(global.DEFAULT_PWD)
	if err != nil {
		return err
	}

	info := &model.User{
		Username: u.Username,
		NickName: u.NickName,
		Avatar:   u.Avatar,
		Sex:      u.Sex,
		Mobile:   u.Mobile,
		Email:    u.Email,
		Remark:   u.Remark,
		Password: pwd,
		UserRole: &model.UserRole{
			RoleCode: u.RoleCode,
		},
	}

	return service.NewUser().AddUser(info)
}
