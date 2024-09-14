package util

import (
	"math/rand"
	"time"
)

type RandomStringGenerator interface {
	GenerateRandomString(length int) string
}

type randomStringGenerator struct {
}

func NewRandomStringGenerator() RandomStringGenerator {
	return randomStringGenerator{}
}

func (randomStringGenerator) GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var result []byte
	for i := 0; i < length; i++ {
		randomIndex := r.Intn(length)
		result = append(result, charset[randomIndex])
	}

	return string(result)

}
