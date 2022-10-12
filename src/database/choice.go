package database

import (
	"fmt"
	"os"
	"strings"
)

func choice() {
	file, err := os.ReadFile("mots.txt")
	if err != nil {
		panic(err)
	}
	line := strings.Split(string(file), "\n")
	fmt.Print(len(line))
}
