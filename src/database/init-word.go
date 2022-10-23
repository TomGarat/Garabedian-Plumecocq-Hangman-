package database

import (
	"bufio"
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
	verbose      bool
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
	return line[x]
}
