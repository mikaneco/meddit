package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func CheckPasswordHash(password string, hash string) bool {
	h := md5.New()
	h.Write([]byte(password))
	passwordHash := hex.EncodeToString(h.Sum(nil))
	return passwordHash == hash
}
