package database

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var maxTries = 10

func (hang *Hangman) getGuess() string {
	reader := bufio.NewReader(os.Stdin)
	guess, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	guess = strings.TrimSpace(guess)
	if guess == "" {
		fmt.Println("please enter a letter")
		return guess
	}
	if strings.Contains(strings.Join(hang.try, ""), guess) {
		fmt.Println("You've already guesses that letter. Please select another one.")
		return ""
	}
	if hang.verbose {
		fmt.Printf("This guess %v \n", guess)
	}
	hang.numTries++
	hang.try = append(hang.try, guess)
	return guess
}

func (hang *Hangman) isMatch(guess string) bool {
	if strings.Contains(hang.word, guess) {
		if hang.verbose {
			fmt.Printf("%v is a match for %v word \n", guess, hang.word)
		}
		return true
	}
	if hang.verbose {
		fmt.Printf("%v is NOT a match for %v word \n", guess, hang.word)
	}
	return false
}

func (hang *Hangman) updateWordState(letter string) {
	if letter == " " {
		for i := 0; i < len(hang.word); i++ {
			hang.wordStatus = append(hang.wordStatus, "_")
		}
	} else {
		for i, l := range hang.word {
			if letter == string(l) {
				hang.wordStatus[i] = letter
			}
		}
	}
}

func (hang *Hangman) continueGame() bool {
	if len(hang.try) == hang.maxTry {
		fmt.Println("you've finished your hangman game, losing.")
		return false
	}
	if strings.Join(hang.wordStatus, "") == hang.word {
		fmt.Println("you've finished your hangman game, and you've won, congrats!")
		return false
	}
	return true
}

func Play() {

	verbose := flag.Bool("v", false, "verbose mode for debugging purposes")
	flag.Parse()

	fmt.Printf(`Welcome to Hangman! I will choose a random word, and you will guess letters you think the word contains.
    You have %v guesses`, maxTries)

	game := Hangman{
		word:    GetWord(),
		maxTry:  maxTries,
		verbose: *verbose,
	}
	game.updateWordState(" ")
	if game.verbose {
		fmt.Printf("Word: %v \n", game.word)
	}

	for game.continueGame() {
		game.DrawBoard()
		guess := game.getGuess()
		if game.isMatch(guess) {
			game.updateWordState(guess)
		}
	}
}
