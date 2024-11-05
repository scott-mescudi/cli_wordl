package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const (
	Red    = "\033[1;31m"
	Green  = "\033[1;32m"
	Magenta = "\033[1;35m"
	Reset  = "\033[0m"

	tr = "┐"
	tl = "┌"

	br = "┘"
	bl = "└"

	vl = "│"
	hl = "─"

)

func clearTerminal() {
	var cmd *exec.Cmd

	
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls") 
	default:
		cmd = exec.Command("clear") 
	}

	cmd.Stdout = os.Stdout 
	cmd.Run()
}


func getRandomWord() (string, error) {
	file, err := os.Open("words.txt")
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	x := rand.Intn(3088)

	var (
		count int
		word  string
	)

	for scanner.Scan() {
		if count == x {
			word = scanner.Text()
			break
		}
		count++
	}

	return word, err
}

func help() {
	box := "██"
	fmt.Printf("%v%v%v = letter is in right spot\n%v%v%v = letter is in word but not in right spot\n%v%v%v = letter is not in word\ntype ':q' to exit\n\n", Green, box, Reset, Magenta, box, Reset, Red, box, Reset)

}

func verifyGuess(guess, word string) (bool, string){
	if guess == word {
		return true, fmt.Sprintf("%v%v%v", Green, word, Reset)
	}

	var builder strings.Builder
	builder.WriteString("-")
	for i := 0; i < len(word); i++ {
		c1 := word[i]
		c2 := guess[i]

		if c1 == c2 {
			builder.WriteString(fmt.Sprintf("%v%s%v-", Green, string(c2), Reset))
		} else if strings.Contains(word, string(c2)) {
			builder.WriteString(fmt.Sprintf("%v%s%v-", Magenta, string(c2), Reset))
		} else {
			builder.WriteString(fmt.Sprintf("%v%s%v-", Red, string(c2), Reset))
		}
	}
	return false, builder.String()
}

func buildTop() string {
	return fmt.Sprintf("%v───────────%v\n", tl, tr)
}

func buildBottom() string {
	return fmt.Sprintf("%v───────────%v", bl, br)
}

func gameBoard(userGuesses []string) string {
	var builder strings.Builder
	builder.WriteString(buildTop())
	
	for _, v := range userGuesses {
		builder.WriteString(fmt.Sprintf("│%v│\n", v))
	}


	builder.WriteString(buildBottom())
	return builder.String()
}

func gameStart() {
	clearTerminal()
	help()
	fmt.Println(gameBoard(nil))
}

func main() {
	word, err := getRandomWord()
	if err != nil {
		log.Fatal(err)
	}

	userGuesses := []string{}
	var guess string
	gameStart()


	for {
		fmt.Print("> ")
		fmt.Scan(&guess)
		if guess == ":q" {
			fmt.Printf("\nThe word was '%v%v%v'",Green, word, Reset)
			break
		}

		if len(guess) != 5{
			fmt.Printf("%vinvalid guess%v: want 5 have %v letters\n", Red, Reset, len(guess))
			continue
		}
		clearTerminal()

		ok, parsedstring := verifyGuess(guess, word)
		userGuesses = append(userGuesses, parsedstring)
		help()

		fmt.Println(gameBoard(userGuesses))

		if ok {
			break
		}
	}
}