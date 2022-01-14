package util

import (
	"math/rand"
	"time"
)

// 随机生成字符串
func RandomString(n int) string {
	var letters = []byte("asdfghjklzxcvbnmqwertyuiop")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
