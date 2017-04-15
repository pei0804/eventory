package utility

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"
)

func ToHash(str string) string {
	converted := sha256.Sum256([]byte(str))
	return hex.EncodeToString(converted[:])
}

func CreateToken(str string) string {
	return ToHash(fmt.Sprintf("%s%s", time.Now().String(), str))
}

func toHashFromScrypt(pass string) (string, error) {
	salt := []byte("some salt")
	converted, err := scrypt.Key([]byte(pass), salt, 16384, 8, 1, 32)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(converted[:]), nil
}

func toHashFromBcrypt(pass string) (string, error) {
	converted, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(converted[:]), nil
}
