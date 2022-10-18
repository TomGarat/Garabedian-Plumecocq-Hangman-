package database

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"
)

func Play() {
	playGame()
}

func init() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go onCtrlC(sigs, done)
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
}

func playGame() {
	word := strings.ToUpper(getWord())
	letters := []byte(word)

	var guessed []byte
	var message string = "How's it hanging?"
	var status int = 0

	for {
		clear()
		banner(false)
		hangman(message, status)

		// Lose state (Hangman is complete)
		if status >= 10 {
			message = "RIP, dude."
			restart(message, status, letters)
		}
		spaces(word + word) // Bubblegum solution: Double the word's length to match word length in restart().
		//Loop through answered letters
		fmt.Print("                 ")
		var hidden []byte
		for i := 0; i < len(letters); i++ {
			if hasKey(guessed, letters[i]) {
				fmt.Print(strings.ToUpper(string(letters[i])) + " ") //Show guessed letters
			} else if letters[i] == '-' {
				guessed = append(guessed, '-') //Show lines in words.
				fmt.Print("- ")
			} else {
				fmt.Print("_ ")
				hidden = append(hidden, '_') //Hide unknown letters
			}
		}

		// Win state (No hidden letters left)
		if !hasKey(hidden, '_') {
			message = "You survived. Play again?"
			restart(message, status, letters)
		}
		fmt.Println()
		fmt.Println()
		fmt.Println("                       ( Press key )")
		fmt.Println()
		fmt.Println()

		//Variables for checking input
		var keys []byte
		isLetter := false
		isNumber := false
		keys = make([]byte, 1)
		os.Stdin.Read(keys)
		key := keys[0]

		//Get key category for later checking (For readability of code)
		switch {
		case key >= 'a' && key <= 'z' || key >= 'A' && key <= 'Z':
			isLetter = true
		case key >= '0' && key <= '9':
			isNumber = true
		case key == 27:
			Play()
		default:
		}

		guess := strings.ToUpper(string(key))

		//Check key press and act accordingly.
		switch {
		case hasKey(letters, key) && hasKey(guessed, key) && isLetter:
			message = "You already got \"" + guess + "\", bro."
			status++
		case hasKey(letters, key) && !hasKey(guessed, key) && isLetter:
			message = "Yep! It has \"" + guess + "\"."
			guessed = append(guessed, key)
		case !isLetter && !isNumber:
			message = "Not a letter!"
		case isNumber:
			message = "That's... a number..."
		default:
			message = "Nope! No \"" + guess + "\"."
			status++
		}
	}
}

func restart(message string, status, diff int, letters []byte) {
	clear()
	banner(false)
	hangman(message, status)

	// Separate letters by space and show the word.
	var word string
	for i := 0; i < len(letters); i++ {
		word += strings.ToUpper(string(letters[i])) + " "
	}

	fmt.Print("                 ")
	spaces(word)
	fmt.Print(word)

	fmt.Println("\n")
	fmt.Print("                         ( Press R to restart ) \n")
	fmt.Print("                         ( Press Esc or Q to go to menu ) ")

	for {
		var keys []byte
		keys = make([]byte, 1)
		os.Stdin.Read(keys)
		key := keys[0]
		switch {
		case key == 'R' || key == 'r':
			status = 0
			message = ""
			playGame()
		case key == 27 || key == 'Q' || key == 'q':
			Play()
		default:
		}
	}
}
func getWord() string {
	wordfile, err := os.Open("database/ressource/words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer wordfile.Close()
	scanner := bufio.NewScanner(wordfile)
	scanner.Split(bufio.ScanLines)
	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(len(words))
	fmt.Println(words[x])
	return words[x]
}

func banner(menu bool) {
	fmt.Println()
	fmt.Println()
	file, err := os.ReadFile("database/ressource/banner.txt")
	if err != nil {
		log.Fatal(err)
	}
	line := strings.Split(string(file), "/n")
	for i := 0; i < len(line)-1; i++ {
		fmt.Print(string(line[i]))
	}
}

func hangman(message string, status int) {
	fmt.Println()
	fmt.Println()
	file, err := os.ReadFile("database/ressource/hangman.txt")
	if err != nil {
		panic(err)
	}
	line := strings.Split(string(file), "/n")
	switch {
	case status == 0:
		fmt.Println()
		for i := 0; i < 6; i++ {
			fmt.Println(line[i])
		}
	case status == 1:
		fmt.Println()
		for i := 7; i < 15; i++ {
			fmt.Println(line[i])
		}
	case status == 2:
		fmt.Println()
		for i := 16; i < 26; i++ {
			fmt.Println(line[i])
		}
	case status == 3:
		fmt.Println()
		for i := 27; i < 37; i++ {
			fmt.Println(line[i])
		}
	case status == 4:
		fmt.Println()
		for i := 38; i < 48; i++ {
			fmt.Println(line[i])
		}
	case status == 5:
		fmt.Println()
		for i := 49; i < 59; i++ {
			fmt.Println(line[i])
		}
	case status == 6:
		fmt.Println()
		for i := 60; i < 70; i++ {
			fmt.Println(line[i])
		}
	case status == 7:
		fmt.Println()
		for i := 71; i < 81; i++ {
			fmt.Println(line[i])
		}
	case status == 8:
		fmt.Println()
		for i := 82; i < 92; i++ {
			fmt.Println(line[i])
		}
	case status == 9:
		fmt.Println()
		for i := 93; i < 103; i++ {
			fmt.Println(line[i])
		}
	case status == 10:
		fmt.Println()
		for i := 104; i < 114; i++ {
			fmt.Println(line[i])
		}
	}
	fmt.Println()
}

func hasKey(slice []byte, key byte) bool {
	for _, a := range slice {
		if strings.ToUpper(string(a)) == strings.ToUpper(string(key)) {
			return true
		}
	}
	return false
}

func spaces(msg string) {
	space := 27 - len(msg)
	for i := 0; i < space; i++ {
		fmt.Print(" ")
	}
}

func onCtrlC(sigs chan os.Signal, done chan bool) {
	sig := <-sigs
	fmt.Println()
	fmt.Println(sig)
	done <- true
	cleanRun()
	clear()
	os.Exit(1)
}

func escape() {
	fmt.Println("Exiting...")
	cleanRun()
	clear()
	os.Exit(0)
}

func cleanRun() {
	exec.Command("stty", "-F", "/dev/tty", "echo").Run()
}

func clear() {
	goos := runtime.GOOS
	switch {
	case goos == "linux" || goos == "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case goos == "windows":
		cmd := exec.Command("cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("Screen clear not supported on your OS.")
		fmt.Println("Please contact author.")
		fmt.Println()
	}
}
