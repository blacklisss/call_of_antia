package util

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"time"
)

func DirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func MakeDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func CreateExpireDate(expireTime time.Duration) time.Time {
	now := time.Now()
	return now.Add(expireTime)
}

func CompareDate(date1, date2 time.Time) bool {
	return date1.Before(date2) || date1.Equal(date2)
}
