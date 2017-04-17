package utility

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

func ToHash(str string) string {
	converted := sha256.Sum256([]byte(str))
	return hex.EncodeToString(converted[:])
}

func CreateToken(str string) string {
	return ToHash(fmt.Sprintf("%s%s", time.Now().String(), str))
}
