package logic

import (
	"encoding/json"
	"errors"
	"fmt"
	"lime/internal/app/admin/model"
	"lime/internal/app/admin/requests"
	"lime/internal/app/admin/response"
	"lime/internal/app/admin/service"
	"lime/internal/common"
	"lime/internal/global"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func FindUserByName(name string) (*response.UserInfo, error) {
	user, err := service.NewUser().GetUserByName(name)
	if err != nil {
		return nil, err
	}

	user.Password = ""

	menuList, err := NewMenuLogic().GetMenuTree(user.UserRole.RoleCode)
	if err != nil {
		return nil, err
	}

	return &response.UserInfo{
		User:  *user,
		Menus: menuList,
	}, nil
}

func Login(info requests.LoginRequest) (model.User, string, error) {
	// 从数据库中查找用户，这里省略数据库操作代码

	userDB, err := service.NewUser().GetUserByName(info.Username)
	if err != nil {
		return model.User{}, "", err
	}

	if userDB.Username != info.Username {
		return model.User{}, "", errors.New("账号密码错误")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(info.Password)); err != nil {
		return model.User{}, "", errors.New("账号密码错误")
	}

	signedToken, err := common.MakeToken(userDB.ID, userDB.UserRole.ID, userDB.Username)
	if err != nil {
		return model.User{}, "", err
	}

	return *userDB, signedToken, nil
}

func GetUser(c *gin.Context) (model.User, error) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		return model.User{}, errors.New("未提供令牌")
	}

	token, err := common.ParseToken(tokenString)
	if err != nil {
		return model.User{}, err
	}

	user := model.User{}
	user.ID = token.ID
	user.Username = token.Username
	user.UserRole = new(model.UserRole)
	user.UserRole.ID = token.RoleID

	data, _ := json.Marshal(&user)
	fmt.Printf("data: %s\n", string(data))
	return user, nil
}

func MakePassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func InitUserRole() {
	userRole := []*model.UserRole{
		{UserID: 1, RoleCode: "9001"},
	}

	for _, ur := range userRole {
		URIsNotExistAdd(ur)
	}
}

func URIsNotExistAdd(ur *model.UserRole) {
	RoleSrv := service.NewUser()
	isHave, err := RoleSrv.CheckUserRole(int(ur.UserID), ur.RoleCode)
	if err != nil {
		global.Logger.Sugar().Errorf("查询用户角色失败，失败原因：%s", err.Error())
		return
	}

	if !isHave {
		err = RoleSrv.AddUserRole(ur)
		if err != nil {
			global.Logger.Sugar().Errorf("添加用户角色失败，失败原因：%s", err.Error())
			return
		}
	}
}

func InitRoleMenus() {
	rms := []*model.RoleMenu{
		{RoleCode: "9001", MenuCode: "9000"},
		{RoleCode: "9001", MenuCode: "9001"},
		{RoleCode: "9001", MenuCode: "9002"},
		{RoleCode: "9001", MenuCode: "9003"},
		{RoleCode: "9001", MenuCode: "9004"},
	}

	for _, rm := range rms {
		RMIsNotExistAdd(rm)
	}
}

func RMIsNotExistAdd(rm *model.RoleMenu) {
	RoleSrv := service.NewUser()
	isHave, err := RoleSrv.CheckRoleMenu(rm.RoleCode, rm.MenuCode)
	if err != nil {
		global.Logger.Sugar().Errorf("查询角色菜单失败，失败原因：%s", err.Error())
		return
	}

	if !isHave {
		err = RoleSrv.AddRoleMenu(rm)
		if err != nil {
			global.Logger.Sugar().Errorf("添加角色菜单失败，失败原因：%s", err.Error())
			return
		}
	}
}
