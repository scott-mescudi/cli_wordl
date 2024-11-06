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

	"github.com/BurntSushi/toml" 
)

var (
	theme       *t.Theme
	letterCount int = 5
	themeName   string = "wordl"
)


type Settings struct {
	Theme       string `toml:"theme"`
	LetterCount int    `toml:"letter_count"`
}

func loadSettings() (*Settings, error) {
	var settings Settings

	
	_, err := toml.DecodeFile("config.toml", &settings)
	if err != nil {
		
		if os.IsNotExist(err) {
			log.Println("config.toml not found, creating a new one with default settings.")
			settings = Settings{
				Theme:       "wordl",  
				LetterCount: 5,        
			}
		
			err := saveSettings(&settings)
			if err != nil {
				return nil, fmt.Errorf("error creating default config.toml: %v", err)
			}
			log.Println("config.toml created successfully.")
			return &settings, nil
		}
		return nil, err
	}
	return &settings, nil
}



func saveSettings(settings *Settings) error {
	file, err := os.Create("config.toml")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := toml.NewEncoder(file)
	err = encoder.Encode(settings)
	if err != nil {
		return err
	}

	return nil
}

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

type wordlists struct {
	Path      string
	WordCount int
}

var wordls = map[int]wordlists{
	5: wordlists{Path: "wordlists/5_letter_words.txt", WordCount: 3088},
	2: wordlists{Path: "wordlists/2_letter_words.txt", WordCount: 88},
	3: wordlists{Path: "wordlists/3_letter_words.txt", WordCount: 682},
	7: wordlists{Path: "wordlists/7_letter_words.txt", WordCount: 4268},
	10: wordlists{Path: "wordlists/10_letter_words.txt", WordCount: 2019},
}

func getRandomWord(lettercount int) (string, error) {
	wl := wordls[lettercount]

	file, err := os.Open(wl.Path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	x := rand.Intn(wl.WordCount)

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
	fmt.Printf("%v%v%v = letter is in right spot\n%v%v%v = letter is in word but not in right spot\n%v%v%v = letter is not in word\ntype ':q' to exit or type ':settings' to change settings\n\n", theme.Colors.AcceptColor, box, t.Reset, theme.Colors.InWordColor, box, t.Reset, theme.Colors.WrongColor, box, t.Reset)
}

func verifyGuess(guess, word string) (bool, string) {
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
	for range letterCount + letterCount + 1 {
		straight += theme.Symbols.Horizontal
	}

	return fmt.Sprintf("%v%s%s%s%v\n", theme.Colors.BoardColor, theme.Symbols.TopLeft, straight, theme.Symbols.TopRight, t.Reset)
}

func BuildBottom(letterCount int) string {
	var straight string
	for range letterCount + letterCount + 1 {
		straight += theme.Symbols.Horizontal
	}
	return fmt.Sprintf("%v%s%s%s%v", theme.Colors.BoardColor, theme.Symbols.BottomLeft, straight, theme.Symbols.BottomRight, t.Reset)
}

func gameBoard(userGuesses []string) string {
	var builder strings.Builder
	builder.WriteString(BuildTop(letterCount)) 

	
	sides := fmt.Sprintf("%v%v%v", theme.Colors.BoardColor, theme.Symbols.Vertical, t.Reset)

	for _, v := range userGuesses {
		
		paddedGuess := v
		for len(paddedGuess) < letterCount {
			paddedGuess += " " 
		}

		builder.WriteString(fmt.Sprintf("%v%s%v\n", sides, paddedGuess, sides))
	}

	builder.WriteString(BuildBottom(letterCount)) 
	return builder.String()
}

func setSettings() bool {
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
	fmt.Print("\nEnter a theme: ")
	fmt.Scan(&t)
	fmt.Print("Enter a letter count: ")
	fmt.Scan(&x)

	if t != "" || t != "default" {
		themeName = t
	}

	if x == "" {
		return false
	}

	xint, err := strconv.Atoi(x)
	if err != nil {
		log.Println("Failed to change letter count")
		return false
	}

	if xint != 0 {
		letterCount = xint
		return true
	}

	return false
}

func gameStart() {
	clearTerminal()
	help()
	fmt.Println(gameBoard(nil))
}

func main() {
	
	settings, err := loadSettings()
	if err != nil {
		log.Fatal(err)
	}

	
	themeName = settings.Theme
	letterCount = settings.LetterCount
	word, err := getRandomWord(letterCount)
	if err != nil {
		log.Fatal(err)
	}

	userGuesses := []string{}
	var guess string

start:
	theme = t.SetTheme(themeName)
	gameStart()

	for {
		theme = t.SetTheme(themeName)
		fmt.Print("> ")
		fmt.Scan(&guess)
		if guess == ":q" {
			fmt.Printf("\nThe word was '%v%v%v'", theme.Colors.AcceptColor, word, t.Reset)
			break
		}

		if guess == ":settings" {
			clearTerminal()
			if ok := setSettings(); ok {
				settings.Theme = themeName
				settings.LetterCount = letterCount
				err := saveSettings(settings)
				if err != nil {
					log.Fatal("Error saving settings:", err)
				}
				fmt.Println("Settings updated! Closing the game...")
				os.Exit(0) 
			}
			goto start
		}

		if len(guess) != letterCount {
			fmt.Printf("%vinvalid guess%v: want %v have %v letters\n", theme.Colors.WrongColor, t.Reset, letterCount, len(guess))
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