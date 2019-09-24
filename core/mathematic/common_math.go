package mathematic

import (
	"encoding/hex"
	"crypto/sha256"
	"time"
	"strconv"
)

func GetHash(hashString string) string {
	hash := sha256.New()
	hash.Write([]byte( strconv.Itoa(int(time.Now().Unix())) + hashString))
	sha256Encode := hex.EncodeToString(hash.Sum(nil))
	return sha256Encode
}