package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func ShaString(str string) string {
	sha := sha256.New()
	sha.Write([]byte(str))
	return hex.EncodeToString(sha.Sum(nil))
}
