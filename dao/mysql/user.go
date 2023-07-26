package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"example/fundemo01/web-app/models"
	"fmt"
)

const secret = "pass1234"

func CheckUserExist(username string) (err error){
	var count int64
	err = Mysql.Raw(`select count(user_id) from user where username = ?`,username).Count(&count).Error
	if err != nil{
		fmt.Println(err)
	}
	if count > 0 {
		return errors.New("User exist!")
	}else{
		fmt.Printf("count is: %d\n",count)
		fmt.Println("User not exist!")
		return nil
	}
}

// InsertUser 向数据库插入一条新的用户数据
func InsertUser(user *models.User){
	user.Password = encryptPassword(user.Password)
	Mysql.Exec(`insert into  user(user_id,username,password) values(?,?,?)`,user.UserID,user.Username,user.Password)
}

// EncryptPassword加密用户密码
func encryptPassword(oPassword string) string{
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Login(user *models.User) (err error){
	enPassword := encryptPassword(user.Password)
	var password string
	rows,err := Mysql.Raw(`select password from user where username = ?`,user.Username).Rows()
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&password); err != nil {
			return err
		}
		if password != enPassword {
			fmt.Println("Password incorrect!")
			return errors.New("User or Password incorrect!")
		} else {
			fmt.Println("Login success!")
			return nil
		}
	}
	return errors.New("User or Password not exist!")
}