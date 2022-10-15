package database

import (
	"math/rand"
	"time"
)

func RandNumber(i int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(i)
}
