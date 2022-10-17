package database

import (
	"fmt"
	"os"
	"strings"
	"text/scanner"
)

// func Choice() {
// 	file, err := os.ReadFile("database/ressource/mots.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	line := strings.Split(string(file), "\n")
// 	i := RandNumber(len(line))
// 	m.Word = append(m.Word, line[i])
//}

func containsAny(s string, chars []string) bool {
	for _, r := range s {
		for _, c := range chars {
			if string(r) == c {
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
	lastIdx := len(s) - 1
	for _, v := range s[0:lastIdx] {
		s += string(v) + sep
	}
	return s + text[lastIdx]
}

func getLtter(found []string) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	for true {
		letter , err := prompt("Enter une lettre: ", Join(found, " "))
		if len(letter) == 1 && containsAny(alphabet,  []string{letter}) {
			return letter
		fmt.Println("Please enter a single letter from the alphabet.")
	}
	return ""
}

func prompt(vals ...interface{}) (string, error) {
	if len(vals) != 0 {
		fmt.Println(vals...)
	}
	sacanner := bufio.NewScanner(os.Stdin)
	sacanner.Scan()
	err := sacanner.Err()
	if err != nil {
		return "", err
	}
	return scanner.Text(), nil
}


func updateFound(found []string, word []string, letter string) []string {
	complete := true
	for i, v := range Word {
		if v == letter {
			found[i] = letter
			complete = false
		}
	}
	return complete
}
