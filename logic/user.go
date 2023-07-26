package logic

import (
	"fmt"
	"go.uber.org/zap"
	"multiovirt-admin/dao/mysql"
	"multiovirt-admin/models"
	"multiovirt-admin/pkg/jwt"
	"multiovirt-admin/pkg/snowflake"
	"multiovirt-admin/settings"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 1、判断用户存不存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
		fmt.Println("User exist!")
		return err
	}
	//2、生成UID
	userID, err := snowflake.GenID(snowflake.SF)
	if err != nil {
		fmt.Printf("snowflake error: %s\n", err)
	} else {
		fmt.Println("snowflake OK!")
	}
	fmt.Println(userID)
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	//3、保存进数据库
	mysql.InsertUser(user)
	return nil
}

func Login(p *models.ParamLogin) (atoken string, rtoken string, err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	if err = mysql.Login(user); err != nil {
		zap.L().Error("database login error:", zap.Error(err))
		return "", "", err
	}
	atoken, rtoken, err = jwt.GenToken(user.UserID, user.Username, settings.Conf.AuthConfig)
	return atoken, rtoken, nil
	//return nil
}
