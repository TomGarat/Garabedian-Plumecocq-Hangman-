package database

import (
	"fmt"
	"os"
	"strings"
)

func Choice() {
	file, err := os.ReadFile("database/ressource/mots.txt")
	if err != nil {
		panic(err)
	}
	line := strings.Split(string(file), "\n")
	i := RandNumber(len(line))
	m.Word = append(m.Word, line[i])
}

func containsAny(s []string, chars string) bool {
	for _, r := range s {
		for _, c := range chars {
			if string(c) == r {
				return true
			}
		}
	}
	return false
}

func Join(text []string, sep string) string {
	if len(text) == 0 {
		return ""
	}
	s := ""
	last := len(s) - 1
	for _, v := range s[0:last] {
		s += string(v) + sep
	}
	return s + text[last]
}

func gatLtter(Word []string) []string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	letters := []string{}
	for true {
		letters = prompt("Enter a letter: ", join(Word, " "))
		if len(letters) == 1 && containsAny(letters, alphabet) {
			return letters
		}
		fmt.Println("Please enter a single letter from the alphabet.")
	}
	return letters
}

func prompt(prompt string, word string) []string {
	fmt.Println(word)
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input)
	return strings.Split(input, "")
}

func updateFound(Word []string, found []string, letter string) []string {
	complete := true
	for i, v := range Word {
		if v == letter {
			found[i] = letter
			complete = false
		}
	}
	return complete
}
