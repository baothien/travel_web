package randutil

import (
	"github.com/google/uuid"
	"gitlab.com/virtual-travel/travel-go-backend/infrastructure/logger"
	"go.uber.org/zap"
	"math/rand"
	"strings"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const number = "0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func UUIDRand() string {
	uuidStr, err := uuid.NewUUID()
	if err != nil {
		logger.Error("Error generate uuid", zap.Error(err))
	}
	return uuidStr.String()
}

//
//// RandomInt generates a random integer between min and max
//func RandomInt(min, max int64) int64 {
//	return min + rand.Int63n(max-min+1)
//}
//
//// RandomString generates a random string of length n
//func RandomString(n int) string {
//	var sb strings.Builder
//	k := len(alphabet)
//
//	for i := 0; i < n; i++ {
//		c := alphabet[rand.Intn(k)]
//		sb.WriteByte(c)
//	}
//
//	return sb.String()
//}

func RandomStringNumber(n int) string {
	var sb strings.Builder
	k := len(number)

	for i := 0; i < n; i++ {
		c := number[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomPhone() string {
	return RandomStringNumber(10)
}

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandomString(length int) string {
	return StringWithCharset(length, charset)
}

func UniqueArrStr(input []string) []string {
	unique := make(map[string]bool)
	var result []string

	for _, s := range input {
		if !unique[s] {
			unique[s] = true
			result = append(result, s)
		}
	}

	return result
}
