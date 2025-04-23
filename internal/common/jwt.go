package common

import (
	"errors"
	"lime/internal/common/config"
	"lime/internal/global"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type TokenInfo struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	RoleCode string `json:"role_code"`
	Token    string `json:"token"`
}

func MakeToken(userID string, roleCode string, name string) (string, error) {
	// 生成 JWT 令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   userID,
		"role_code": roleCode,
		"username":  name,
		"exp":       time.Now().Add(time.Hour * 24).Unix(), // 设置过期时间，这里示例为1天
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

func ParseToken(tokenString string) (TokenInfo, error) {
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
		return TokenInfo{}, err
	}

	if !token.Valid {
		return TokenInfo{}, errors.New("无效的令牌")
	}

	// 获取用户 ID 和用户名
	mc := token.Claims.(jwt.MapClaims)
	userID := mc["user_id"].(string)
	roleCode := mc["role_code"].(string)
	username := mc["username"].(string)

	return TokenInfo{UserID: userID, RoleCode: roleCode, Username: username}, nil
}

func RefreshToken(oldToken string) (TokenInfo, error) {
	token, err := ParseToken(oldToken)
	if err != nil {
		return TokenInfo{}, err
	}

	// 生成一个新的 JWT 令牌
	signedToken, err := MakeToken(token.UserID, token.RoleCode, token.Username)
	if err != nil {
		return TokenInfo{}, err
	}

	token.Token = signedToken
	return token, nil
}

func MakePassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
