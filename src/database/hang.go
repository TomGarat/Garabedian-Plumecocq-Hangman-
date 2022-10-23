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
	if len(guess) > 1 {
		PrintSlowl("Vous devez entrer une seule lettre !\n", 1)
		return ""
	}
	if strings.Contains("12345678910", guess) {
		PrintSlowl("Vous devez entrer une lettre !\n", 1)
		return ""
	}
	if strings.Contains("ABCDEFGHIJKLMNOPQRSTUVWXYZ", guess) {
		PrintSlowl("Vous devez entrer une lettre minuscule !\n", 1)
		return ""
	}
	if strings.Contains("&é'(-è_çà)=°~#{[|`@]}$£¤µ%ù*+§/!.,:;?<>", guess) {
		PrintSlowl("Vous devez entrer une lettre !\n", 1)
		return ""
	}
	hang.lettersTried = append(hang.lettersTried, guess)
	return guess
}

func (hang *Hangman) updateWordState(letter string) {
	if letter == " " {
		for i := 0; i < len(hang.word); i++ {
			hang.wordStatus = append(hang.wordStatus, "_")
		}
	} else {
		for i, l := range hang.word {
			if string(l) == letter {
				hang.wordStatus[i] = letter
			}
		}
		if !strings.Contains(hang.word, letter) {
			hang.numTries++
			PrintSlowl(hang.print[hang.numTries], 1)
		}
	}
}

func (hang *Hangman) continueGame() bool {
	if hang.numTries >= hang.maxTry {
		PrintSlowl("Vous avez perdu !\n", 1)
		return false
	}
	if strings.Join(hang.wordStatus, "") == hang.word {
		PrintSlowl("Vous avez gagné !\n", 1)
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
		successful:   false,
	}
	game.updateWordState(" ")

	for game.continueGame() {
		game.DrawBoard()
		guess := game.getGuess()
		game.updateWordState(guess)
	}
}
