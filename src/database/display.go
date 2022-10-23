package database

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func (hang *Hangman) DrawBoard() {
	tableaux := []string{
		fmt.Sprint("vous avez ", hang.maxTry-hang.numTries, " essais"),
		fmt.Sprint("vaus avez deja essaye : ", hang.lettersTried),
		fmt.Sprint("le mot est : ", hang.wordStatus),
	}
	for _, line := range tableaux {
		PrintSlowl(line+"\n", 1)
		time.Sleep(1 * time.Second)
	}
}

func PrintSlowl(text string, delay int) {
	for _, char := range text {
		fmt.Print(string(char))
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
}

func (hang Hangman) DisplayHangman() {
	fmt.Println(hang.lettersTried[10-hang.maxTry-hang.numTries])
}

func ReadHangmansFile() []string {
	var filename string = "database/ressource/hangman.txt"
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error parsing file %s\n", filename)
		fmt.Println(err)
	}
	return hangmansFromBytes(content)
}
func hangmansFromBytes(b []byte) []string {
	var arr []string = make([]string, 0)
	var tmp string
	arr = append(arr, "")
	for index, _byte := range b {
		if index != 0 && index%107 == 0 {
			arr = append(arr, tmp)
			tmp = ""
		}
		tmp += string(_byte)
	}
	arr = append(arr, tmp)
	return arr
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
