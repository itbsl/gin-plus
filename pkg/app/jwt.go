package app

import (
	"gin-plus/global"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// Claims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 如果我们需要记录额外的字段信息，需要自定义结构体
type Claims struct {
	UserId int64 `json:"user_id"`
	jwt.StandardClaims
}

// Generate 生成token
func Generate(UserId int64) (string, error) {
	now := time.Now()
	expiredTime := now.Add(time.Second * global.Config.AppConfig.JWTTokenExpired)
	claims := Claims{
		UserId: UserId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredTime.Unix(),
			Issuer:    global.Config.AppConfig.Name,
		},
	}
	//使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//使用指定的secret签名并获得完整编码后的字符串token
	return token.SignedString([]byte(global.Config.AppConfig.JWTSecret))
}

// Parse 解析并校验token
func Parse(tokenStr string) (*Claims, error) {
	// ParseWithClaims用于解析鉴权的声明，方法内部是具体的解码和校验的过程，最终返回*Token
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.Config.AppConfig.JWTSecret), nil
	})
	if token != nil {
		//校验token
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, err
}
