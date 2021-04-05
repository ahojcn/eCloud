package util

import (
	"crypto/md5"
	"fmt"
)

func Md5Str(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}
