package util

import (
	"golang.org/x/crypto/bcrypt"
)

// IsValidPassword は入力されたパスワードを検証し、
// 有効なパスワードの場合、trueを返します。
func IsValidPassword(p string) bool {
	if len(p) > 72 || len(p) == 0 {
		return false
	}
	for _, v := range p {
		if v < 0x20 || v > 0x7e {
			return false
		}
	}
	return true
}

// PasswordHash は入力されたパスワードをハッシュ化し、
// 文字列として返します。
func PasswordHash(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

// PasswordVerify は入力されたパスワードと、ハッシュ化された
// パスワードを比較し一致した場合、trueを返します。
func PasswordVerify(hash, pass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
}
