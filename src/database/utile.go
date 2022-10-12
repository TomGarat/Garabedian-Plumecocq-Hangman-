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
	rand.Intn(i)
}
