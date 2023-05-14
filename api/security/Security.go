package security

import (
	"crypto/md5"
	"encoding/hex"
)

//Функция хэширования пароля
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
