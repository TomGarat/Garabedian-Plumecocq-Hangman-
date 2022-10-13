package database

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func choice() {
	file, err := os.ReadFile("database/ressource/mots.txt")
	if err != nil {
		panic(err)
	}
	line := strings.Split(string(file), "\n")
	i := RandNumber(len(line))
	fmt.Println(line[i])
	sacanner := bufio.NewScanner(os.Stdin)
	sacanner.Scan(i)
}
