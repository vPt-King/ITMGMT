package utils

import (
	"encoding/hex"
)

func GenerateSessionID() string {
	b := make([]byte, 16) // 128-bit entropy
	return hex.EncodeToString(b)
}
