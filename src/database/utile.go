package database

import (
	"fmt"
	"math/rand"
	"time"
)

func Test() {
	fmt.Println("Hello World!")
}

func RandNumber(i int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(i)
}
