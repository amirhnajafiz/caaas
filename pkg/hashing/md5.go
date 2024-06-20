package hashing

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5Hash is used to hash any string into md5 format.
func MD5Hash(text string) string {
	hash := md5.Sum([]byte(text))

	return hex.EncodeToString(hash[:])
}
