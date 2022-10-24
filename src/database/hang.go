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

func (hang *Hangman) getGuess() string { //cette fonction permet de récupérer la lettre entrée par l'utilisateur et de vérifier qu'elle n'a pas déjà été essayée
	reader := bufio.NewReader(os.Stdin)   //on lit l'entrée de l'utilisateur avec bufio
	guess, err := reader.ReadString('\n') //on lit la lettre entrée par l'utilisateur
	if err != nil {                       //si il y a une erreur on affiche un message d'erreur
		log.Fatal(err)
	}
	guess = strings.TrimSpace(guess)                                  //on enlève les espaces avant et après la lettre
	if strings.Contains(strings.Join(hang.lettersTried, ""), guess) { //si la lettre a déjà été essayée, on affiche un message d'erreur
		PrintSlowl("Vous avez déjà essayé cette lettre !\n", 1)
		return ""
	}
	if len(guess) > 1 { //si l'utilisateur a entré plus d'une lettre, on affiche un message d'erreur
		PrintSlowl("Vous devez entrer une seule lettre !\n", 1)
		return ""
	}
	if !strings.Contains("abcdefghijklmnopqrstuvwxyz", guess) { //si l'utilisateur n'a pas entré une lettre, on affiche un message d'erreur
		PrintSlowl("Vous devez entrer une lettre !\n", 1)
		return ""
	}
	hang.lettersTried = append(hang.lettersTried, guess) //on ajoute la lettre à la liste des lettres essayées
	return guess
}

func (hang *Hangman) updateWordState(letter string) { //cette fonction permet de mettre à jour l'état du mot en fonction de la lettre entrée par l'utilisateur
	if letter == " " {
		for i := 0; i < len(hang.word); i++ {
			hang.wordStatus = append(hang.wordStatus, "_") //on ajoute un "_" pour chaque lettre du mot
		}
	} else {
		for i, l := range hang.word { //pour chaque lettre du mot on vérifie si elle correspond à la lettre entrée par l'utilisateur
			if string(l) == letter { //si oui, on remplace le "_" par la lettre
				hang.wordStatus[i] = letter
			}
		}
		if !strings.Contains(hang.word, letter) { //si la lettre n'est pas dans le mot, on incrémente le ascci art
			hang.numTries++
			PrintSlowl(hang.print[hang.numTries], 1)
		}
	}
}

func (hang *Hangman) continueGame() bool { //cette fonction permet de vérifier si le jeu doit continuer ou non
	if hang.numTries >= hang.maxTry { //si le nombre d'essais est supérieur au nombre d'essais maximum, on affiche un message d'erreur et on retourne false
		PrintSlowl("Vous avez perdu !\n", 1)
		return false
	}
	if strings.Join(hang.wordStatus, "") == hang.word { //si le mot est trouvé, on affiche un message de victoire et on retourne false
		AfficheWin() //on affiche le message de victoire
		return false
	}
	return true
}

func Play() { //cette fonction permet de lancer le jeu

	Banner() //on affiche le banner
	tableaux := []string{
		"Bienvenue dans le jeu du pendu !",
		fmt.Sprintf("Vous avez %v essais pour trouver le mot mystère", maxTries),
		"Bon courage !",
	}
	for _, line := range tableaux { //on affiche les messages
		PrintSlowl(line+"\n", 1)
		time.Sleep(1 * time.Second)
	}

	game := Hangman{ //on initialise les variables du jeu
		wordStatus:   []string{},
		lettersTried: []string{},
		maxTry:       maxTries,
		numTries:     0,
		word:         GetWord(),
		successful:   false,
	}
	game.updateWordState(" ") //on met à jour l'état du mot

	for game.continueGame() { //tant que le jeu doit continuer, on demande à l'utilisateur de rentrer une lettre
		game.DrawBoard()
		guess := game.getGuess()
		game.updateWordState(guess)
	}
}
