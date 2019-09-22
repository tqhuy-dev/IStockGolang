package mathematic

import (
	"encoding/hex"
	"crypto/sha256"
)

func GetHash(hashString string) string {
	hash := sha256.New()
	hash.Write([]byte(hashString))
	sha256Encode := hex.EncodeToString(hash.Sum(nil))
	return sha256Encode
}