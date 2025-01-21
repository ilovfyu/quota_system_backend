package utils

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

const charset = "DEFGHIJKLMNabcdefghijklmnopqrst1234uvwxyzABCOPQRSTUVWXYZ056789"

func GeneratorGuid() string {
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, 32)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func IsValidJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

func HashPassword(password string) (string, error) {
	// 生成哈希
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func VerifyPassword(password, hashPassword string) bool {
	// 检查密码是否匹配
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}
