package logic

import (
	"errors"
	"fmt"
	"lime/internal/app/admin/model"
	"lime/internal/app/admin/requests"
	"lime/internal/app/admin/response"
	"lime/internal/app/admin/service"
	"lime/internal/common"
	"strconv"

	"github.com/gin-gonic/gin"
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

	err = common.ComparePassword(userDB.Password, info.Password)
	if err != nil {
		return model.User{}, "", errors.New("账号密码错误")
	}

	signedToken, err := common.MakeToken(fmt.Sprintf("%d", userDB.ID), userDB.UserRole.RoleCode, userDB.Username)
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
	userid, err := strconv.Atoi(token.UserID)
	if err != nil {
		return model.User{}, err
	}
	user.ID = uint(userid)
	user.Username = token.Username
	user.UserRole = new(model.UserRole)
	user.UserRole.RoleCode = token.RoleCode

	// data, _ := json.Marshal(&user)
	// fmt.Printf("data: %s\n", string(data))
	return user, nil
}
