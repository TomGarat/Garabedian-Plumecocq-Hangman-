package database

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
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
		PrintSlowl("Vous devez entrer une lettre !\n", 1)
		return guess
	}
	if strings.Contains(strings.Join(hang.lettersTried, ""), guess) {
		PrintSlowl("Vous avez déjà essayé cette lettre !\n", 1)
		return ""
	}
	hang.numTries++
	hang.lettersTried = append(hang.lettersTried, guess)
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
	if len(hang.lettersTried) == hang.maxTry {
		PrintSlowl("Vous avez perdu !", 1)
		return false
	}
	if strings.Join(hang.wordStatus, "") == hang.word {
		PrintSlowl("Vous avez gagné !", 1)
		return false
	}
	return true
}

func Play() {

	tableaux := []string{
		"Bienvenue dans le jeu du pendu !",
		fmt.Sprintf("Vous avez %v essais pour trouver le mot mystère", maxTries),
		"Bon courage !",
	}
	for _, line := range tableaux {
		PrintSlowl(line+"\n", 1)
		time.Sleep(1 * time.Second)
	}

	game := Hangman{
		wordStatus:   []string{},
		lettersTried: []string{},
		maxTry:       maxTries,
		numTries:     0,
		word:         GetWord(),
		verbose:      false,
	}
	game.updateWordState(" ")

	for game.continueGame() {
		game.DrawBoard()
		guess := game.getGuess()
		if game.isMatch(guess) {
			game.updateWordState(guess)
		}
	}
}
