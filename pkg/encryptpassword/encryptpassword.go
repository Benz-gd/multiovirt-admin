package encryptpassword

import (
	"crypto/md5"
	"encoding/hex"
)

const secret = "pass1234"


// EncryptPassword加密用户密码
func EncryptPassword(oPassword string) string{
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
