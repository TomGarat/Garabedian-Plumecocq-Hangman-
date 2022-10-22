package database

import (
	"fmt"
	"os"
	"strings"
)

func (hang *Hangman) DrawBoard() {
	fmt.Printf("Guesses left: %v \n", hang.maxTry-hang.numTries)
	fmt.Printf("Guesses: %v \n", hang.try)
	fmt.Println(hang.wordStatus)
	hang.DrawHangman(hang.numTries)

}

func (hang *Hangman) DrawHangman(status int) {
	file, err := os.ReadFile("database/ressource/hangman.txt")
	if err != nil {
		panic(err)
	}
	line := strings.Split(string(file), "/n")
	switch {
	case status == 0:
		fmt.Println()
		for i := 0; i < 6; i++ {
			fmt.Println(line[i])
		}
	case status == 1:
		fmt.Println()
		for i := 7; i < 15; i++ {
			fmt.Println(line[i])
		}
	case status == 2:
		fmt.Println()
		for i := 16; i < 26; i++ {
			fmt.Println(line[i])
		}
	case status == 3:
		fmt.Println()
		for i := 27; i < 37; i++ {
			fmt.Println(line[i])
		}
	case status == 4:
		fmt.Println()
		for i := 38; i < 48; i++ {
			fmt.Println(line[i])
		}
	case status == 5:
		fmt.Println()
		for i := 49; i < 59; i++ {
			fmt.Println(line[i])
		}
	case status == 6:
		fmt.Println()
		for i := 60; i < 70; i++ {
			fmt.Println(line[i])
		}
	case status == 7:
		fmt.Println()
		for i := 71; i < 81; i++ {
			fmt.Println(line[i])
		}
	case status == 8:
		fmt.Println()
		for i := 82; i < 92; i++ {
			fmt.Println(line[i])
		}
	case status == 9:
		fmt.Println()
		for i := 93; i < 103; i++ {
			fmt.Println(line[i])
		}
	case status == 10:
		fmt.Println()
		for i := 104; i < 114; i++ {
			fmt.Println(line[i])
		}
	}
	fmt.Println()
}
