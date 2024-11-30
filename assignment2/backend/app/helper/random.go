package helper

import (
	"math/rand"
	"time"
)

func RandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(10000) + 1
	return randomNumber
}
