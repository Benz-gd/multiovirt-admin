package jwt

import (
	"errors"
	"example/fundemo01/web-app/settings"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"time"
)

//const (
//	aTokenExpireDuration = time.Minute*30
//	rTokenExpireDuration = time.Hour*36
//)

var customSecret = []byte("cGFzc3dvcmQK")
type CustomClaims struct{
	UserID int64 `json:"user_id"`
	UserName string `json:"username"`
	jwt.RegisteredClaims
}

// GenToken :生成Token
func GenToken(userid int64,username string,cfg *settings.Auth) (aToken string, rToken string,err error) {

	aClaims := CustomClaims{
		userid,
		username,
		jwt.RegisteredClaims{
			//ExpiresAt: jwt.NewNumericDate(time.Now().Add(aTokenExpireDuration)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second*time.Duration(cfg.Access_Token_Expire))),
			Issuer: "web-app",
		},
	}
	rClaims := jwt.RegisteredClaims{
		//ExpiresAt: jwt.NewNumericDate(time.Now().Add(rTokenExpireDuration)),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour*time.Duration(cfg.Refresh_Token_Expire))),
		Issuer: "web-app",
	}
	//使用SigningMethodHS256加密方法对claims进行加密
	aToken,err = jwt.NewWithClaims(jwt.SigningMethodHS256,aClaims).SignedString(customSecret)
	if err != nil {
		zap.L().Error("GenToken access token error",zap.Error(err))
	}
	rToken,err = jwt.NewWithClaims(jwt.SigningMethodHS256,rClaims).SignedString(customSecret)
	if err != nil{
		zap.L().Error("GenToken refresh token error",zap.Error(err))
	}
	//使用customSecret进行签名，并返回加密token串
	return aToken,rToken,nil
}
// ParseToken :解析token
func ParseToken(aToken string) (*CustomClaims,error){
	cc := new(CustomClaims)
	token,err := jwt.ParseWithClaims(aToken,cc,func(token *jwt.Token)(i interface{},err error){
		return customSecret,nil
	})
	if err != nil {
		return nil,err
	}
	//对token对象重的claim进行类型断言
	if token.Valid {
		return cc,nil
	}
	return nil,errors.New("invalid token")
}

// RefreshToken刷新AccessToken
func RefreshToken(aToken,rToken string) (newAToken string, err error) {
	//从refresh token中检查token是否有效
	if _,err = jwt.Parse(rToken,func(token *jwt.Token)(i interface{},err error){
		return "",nil
	});err != nil {
		return "",err
	}
	return
}














