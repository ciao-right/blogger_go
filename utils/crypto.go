package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// md5
func Check(content, encrypted string) bool {
	return strings.EqualFold(Encode(content), encrypted)
}

func Encode(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
