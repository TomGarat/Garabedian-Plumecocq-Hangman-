package database

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func (hang *Hangman) DrawBoard() { //cette fonction permet d'afficher les info du jeu
	hang.print = ReadHangmansFile()
	tableaux := []string{
		fmt.Sprint("vous avez ", hang.maxTry-hang.numTries, " essais"),
		fmt.Sprint("vous avez deja essaye : ", hang.lettersTried),
		fmt.Sprint("le mot est : ", hang.wordStatus),
	}
	for _, line := range tableaux {
		PrintSlowl(line+"\n", 1)
		time.Sleep(1 * time.Second)
	}
}
func Banner() { //cette fonction permet d'afficher le logo du jeu
	var filename string = "database/ressource/banner.txt"
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error parsing file %s\n", filename)
		fmt.Println(err)
	}
	PrintSlowl(string(content), 1)
	fmt.Println("")
	fmt.Println("")
}

func AfficheWin() {
	var filename string = "database/ressource/win.txt"
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error parsing file %s\n", filename)
		fmt.Println(err)
	}
	PrintSlowl(string(content), 1)
	fmt.Println("")
	fmt.Println("")
}
func PrintSlowl(text string, delay int) { //cette fonction permet d'afficher le texte caractère par caractère ps c'est pour ethan :)
	for _, char := range text {
		fmt.Print(string(char))
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
}

func ReadHangmansFile() []string { //cette fonction permet de lire le fichier hangmans.txt
	var filename string = "database/ressource/hangman.txt"
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error parsing file %s\n", filename)
		fmt.Println(err)
	}
	return hangmansFromBytes(content)
}
func hangmansFromBytes(b []byte) []string { //cette fonction permet de lire le fichier hangmans.txt et de le mettre dans un tableau de string (une ligne = un string)
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
