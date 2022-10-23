package database

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func (hang *Hangman) DrawBoard() {
	hang.print = ReadHangmansFile()
	tableaux := []string{
		fmt.Sprint("vous avez ", hang.maxTry-hang.numTries, " essais"),
		fmt.Sprint("vaus avez deja essaye : ", hang.lettersTried),
		fmt.Sprint("le mot est : ", hang.wordStatus),
		hang.print[hang.numTries],
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
	lines := strings.Split(string(b), "\n")
	for index, line := range lines {
		if index != 0 && index%10 == 0 {
			arr = append(arr, tmp)
			tmp = ""
		}
		tmp += line + "\n"
	}
	arr = append(arr, tmp)
	return arr
}
