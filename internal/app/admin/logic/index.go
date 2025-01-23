package logic

import (
	"errors"
	"fmt"
	"lime/internal/app/admin/model"
	"lime/internal/app/admin/requests"
	"lime/internal/app/admin/response"
	"lime/internal/app/admin/service"
	"lime/internal/common/config"
	"lime/internal/global"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func FindUserByName(name string) (*response.UserInfo, error) {
	user, err := service.NewUser().GetUserByName(name)
	if err != nil {
		return nil, err
	}

	user.Password = ""

	menuList, err := GetMenuTree(user.UserRole.RoleCode)
	if err != nil {
		return nil, err
	}

	return &response.UserInfo{
		User:  *user,
		Menus: menuList,
	}, nil
}

func GetHomeCount() (map[string]any, error) {
	userCount, err := service.NewUser().GetAllCount()
	if err != nil {
		return nil, err
	}

	return map[string]any{
		"userCount": userCount,
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

	signedToken, err := MakeToken(userDB)
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

	return GetUserByToken(tokenString)
}

func MakeToken(userDB *model.User) (string, error) {
	// 生成 JWT 令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  userDB.ID,
		"username": userDB.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // 设置过期时间，这里示例为1天
	})
	// 这里的 secretKey 应该妥善保管，例如从环境变量中获取等
	key := config.GetParam(config.JWT, "jwt-secret-key", "pkkwmjjum5hvfqybnbxo97ol2spriy49").String()
	secretKey := []byte(key)

	tokenStr, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	// 给 token 添加 Bearer 前缀
	return tokenStr, nil
}

func GetUserByToken(tokenString string) (model.User, error) {
	fmt.Println("tokenString: ", tokenString)
	// 解析旧的 JWT 令牌
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("无效的签名方法")
		}

		key := config.GetParam(config.JWT, "jwt-secret-key", "pkkwmjjum5hvfqybnbxo97ol2spriy49").String()
		return []byte(key), nil
	})

	if err != nil {
		global.Logger.Sugar().Errorf("解析令牌失败: %s", err.Error())
		return model.User{}, err
	}

	if !token.Valid {
		return model.User{}, errors.New("无效的令牌")
	}

	// 获取用户 ID 和用户名
	mc := token.Claims.(jwt.MapClaims)
	userID := mc["user_id"].(float64)
	username := mc["username"].(string)

	u := model.User{}
	u.ID = uint(userID)
	u.Username = username
	return u, nil
}

func RefreshToken(oldToken string) (model.User, string, error) {
	user, err := GetUserByToken(oldToken)
	if err != nil {
		return model.User{}, "", err
	}

	// 生成一个新的 JWT 令牌
	signedToken, err := MakeToken(&user)
	if err != nil {
		return model.User{}, "", err
	}

	return user, signedToken, nil
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
