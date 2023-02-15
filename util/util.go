package util

import "crypto/sha1"

func Sha1(str string) string {
	h := sha1.New()
	return string(h.Sum([]byte(str)))
}