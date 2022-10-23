package database

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var (
	hang Hangman
)

type Hangman struct {
	wordStatus   []string
	lettersTried []string
	maxTry       int
	numTries     int
	word         string
	successful   bool
	print        []string
}

func GetWord() string {
	wordfile, err := os.Open("database/ressource/words.txt")
	if err != nil {
		panic(err)
	}
	defer wordfile.Close()
	scanner := bufio.NewScanner(wordfile)
	scanner.Split(bufio.ScanLines)
	var line []string
	for scanner.Scan() {
		line = append(line, scanner.Text())
	}
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(len(line))
	fmt.Print(line[x])
	return line[x]
}
