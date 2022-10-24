package database

import (
	"fmt"
	"math/rand"
	"time"
)

func Test() {
	fmt.Println("Hello World!")
}

func RandNumber(i int) {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(i)
	fmt.Println(i)
}
func RandNumberMinMax(min, max int) int {
	rand.NewSource(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
