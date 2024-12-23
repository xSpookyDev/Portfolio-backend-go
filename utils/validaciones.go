package utils

import (
	"regexp"
	"unicode"
)

var Regex_correo = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

func ValidarPassword(s string) bool {
	var (
		hasMinLen = len(s) >= 8
		hasUpper  = false
		hasLower  = false
		hasNumber = false
	)
	for _, char := range s {
		switch {
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber
}
