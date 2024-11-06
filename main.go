package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	t "wordl/internal/themes"
)

var (
	theme *t.Theme
	letterCount int = 5
	themeName string = "wordl"
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
	file, err := os.Open("wordlists/5_letter_words.txt")
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
	fmt.Printf("%v%v%v = letter is in right spot\n%v%v%v = letter is in word but not in right spot\n%v%v%v = letter is not in word\ntype ':q' to exit\n\n", theme.Colors.AcceptColor, box, t.Reset, theme.Colors.InWordColor, box, t.Reset, theme.Colors.WrongColor, box, t.Reset)

}

func verifyGuess(guess, word string) (bool, string){
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("%v-%v", theme.Colors.BoardColor, t.Reset))
	for i := 0; i < len(word); i++ {
		c1 := word[i]
		c2 := guess[i]

		if c1 == c2 {
			builder.WriteString(fmt.Sprintf("%v%v%s%v%v-%v", t.Reset, theme.Colors.AcceptColor, string(c2), t.Reset, theme.Colors.BoardColor, t.Reset))
		} else if strings.Contains(word, string(c2)) {
			builder.WriteString(fmt.Sprintf("%v%v%s%v%v-%v", t.Reset, theme.Colors.InWordColor, string(c2), t.Reset, theme.Colors.BoardColor, t.Reset))
		} else {
			builder.WriteString(fmt.Sprintf("%v%v%s%v%v-%v", t.Reset, theme.Colors.WrongColor, string(c2), t.Reset, theme.Colors.BoardColor, t.Reset))
		}
	}
	if guess == word {
		return true, builder.String()
	}

	return false, builder.String()
}

func BuildTop(letterCount int) string {
	var straight string
	for range letterCount + letterCount + 1{
		straight += theme.Symbols.Horizontal
	}

	return fmt.Sprintf("%v%s%s%s%v\n", theme.Colors.BoardColor, theme.Symbols.TopLeft, straight, theme.Symbols.TopRight, t.Reset)
}

func BuildBottom(letterCount int) string {
	var straight string
	for range letterCount + letterCount + 1{
		straight += theme.Symbols.Horizontal
	}
	return fmt.Sprintf("%v%s%s%s%v", theme.Colors.BoardColor, theme.Symbols.BottomLeft,straight,theme.Symbols.BottomRight, t.Reset)
}

func gameBoard(userGuesses []string) string {
	var builder strings.Builder
	builder.WriteString(BuildTop(letterCount))

	var sides = fmt.Sprintf("%v%v%v", theme.Colors.BoardColor, theme.Symbols.Vertical, t.Reset)
	
	for _, v := range userGuesses {
		builder.WriteString(fmt.Sprintf("%[1]v%v%[1]v\n", sides, v))
	}


	builder.WriteString(BuildBottom(letterCount))
	return builder.String()
}

func setSettings() {
	fmt.Println(`
	Themes:
		- wordl     (default)
		- oceanbreeze
		- desertsunset
		- neonnight
		- forestmeadow
		- cyberpunkglow
		- monochromepop
		- impossible

	letter count:
		- 2 letters
		- 3 letters
		- 5 letters (default)
		- 7 letters
		- 10 letters
	`)

	var (
		t string
		x string
	)
	fmt.Println("Enter a theme: ")
	fmt.Scan(&t)
	fmt.Println("\nEnter a letter count: ")
	fmt.Scan(&x)

	if t != "" || t != "default"{
		themeName = t
	}

	if x == ""{
		return
	}
	
	xint, err := strconv.Atoi(x)
	if err != nil {
		log.Println("Failed to change letter count")
		return
	}
	
	if xint != 5 {
		letterCount = xint
	}
	
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

	theme = t.SetTheme(themeName)
	userGuesses := []string{}
	var guess string
	gameStart()

	for {
		theme = t.SetTheme(themeName)
		fmt.Print("> ")
		fmt.Scan(&guess)
		if guess == ":q" {
			fmt.Printf("\nThe word was '%v%v%v'",theme.Colors.AcceptColor, word, t.Reset)
			break
		}

		if guess == ":settings" {
			clearTerminal()
			setSettings()
			clearTerminal()
			continue
		}

		if len(guess) != 5{
			fmt.Printf("%vinvalid guess%v: want 5 have %v letters\n", theme.Colors.WrongColor, t.Reset, len(guess))
			continue
		}
		clearTerminal()

		ok, parsedstring := verifyGuess(guess, word)
		userGuesses = append(userGuesses, parsedstring)
		help()

		fmt.Println(gameBoard(userGuesses))

		if ok {
			fmt.Println("Congratulations!!!")
			break
		}
	}
}