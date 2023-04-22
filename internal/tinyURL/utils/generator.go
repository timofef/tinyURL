package utils

import (
	"math/rand"
	"time"
)

var strLen = 10
var charset = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateString() string {
	randomString := make([]rune, strLen)
	for i := 0; i < len(randomString); i++ {
		randomString[i] = charset[rand.Intn(len(randomString))]
	}

	return string(randomString)
}
