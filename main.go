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



// Reset
const Reset = "\033[0m"

// Default Theme
const (
	Red    = "\033[1;31m"
	Green  = "\033[1;32m"
	Magenta = "\033[1;35m"

	tr = "┐"
	tl = "┌"
	br = "┘"
	bl = "└"
	h  = "─"
	v  = "│"
)

// Ocean Breeze Theme
const (
	ObDeepBlue   = "\033[1;34m"
	ObCyan       = "\033[1;36m"
	ObLightBlue  = "\033[1;94m"
	ObWhite      = "\033[1;37m"

	ObTR = "╮"
	ObTL = "╭"
	ObBR = "╯"
	ObBL = "╰"
	ObH  = "─"
	ObV  = "│"
)

// Desert Sunset Theme
const (
	DsRed    = "\033[1;31m"
	DsOrange = "\033[1;38;5;214m" // RGB escape for warm orange
	DsYellow = "\033[1;33m"
	DsBrown  = "\033[1;38;5;94m"  // RGB escape for brownish color

	DsTR = "▞"
	DsTL = "▚"
	DsBR = "▚"
	DsBL = "▞"
	DsH  = "━"
	DsV  = "┃"
)

// Neon Night Theme
const (
	NnPurple  = "\033[1;35m"
	NnLime    = "\033[1;92m"
	NnCyan    = "\033[1;96m"
	NnMagenta = "\033[1;95m"

	NnTR = "╖"
	NnTL = "╓"
	NnBR = "╜"
	NnBL = "╙"
	NnH  = "═"
	NnV  = "║"
)

// Forest Meadow Theme
const (
	FmDarkGreen  = "\033[1;32m"
	FmLightGreen = "\033[1;92m"
	FmBrown      = "\033[1;33m" // Yellow for a brownish tone
	FmYellow     = "\033[1;93m"

	FmTR = "╮"
	FmTL = "╭"
	FmBR = "╯"
	FmBL = "╰"
	FmH  = "─"
	FmV  = "│"
)

// Cyberpunk Glow Theme
const (
	CgMagenta    = "\033[1;35m"
	CgBrightCyan = "\033[1;96m"
	CgBlue       = "\033[1;34m"
	CgYellow     = "\033[1;33m"

	CgTR = "═╗"
	CgTL = "╔"
	CgBR = "╝"
	CgBL = "╚"
	CgH  = "═"
	CgV  = "║"
)

// Monochrome Pop Theme
const (
	MpWhite     = "\033[1;37m"
	MpGray      = "\033[1;90m"
	MpBlack     = "\033[1;30m"
	MpLightGray = "\033[1;97m"

	MpTR = "»"
	MpTL = "«"
	MpBR = "«"
	MpBL = "»"
	MpH  = "─"
	MpV  = "│"
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
	fmt.Printf("%v%v%v = letter is in right spot\n%v%v%v = letter is in word but not in right spot\n%v%v%v = letter is not in word\ntype ':q' to exit\n\n", Green, box, Reset, Magenta, box, Reset, Red, box, Reset)

}

func verifyGuess(guess, word string) (bool, string){
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
	if guess == word {
		return true, builder.String()
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
			fmt.Println("Congratulations!!!")
			break
		}
	}
}
// TODO: add option for multiple letter words like 7, 6, 3, 20 etc
// TODO: filter dict for n letter words