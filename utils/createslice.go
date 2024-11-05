package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("../wordlists/5_letter_words.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var builder strings.Builder
	builder.WriteString("var words = []string{")
	for scanner.Scan() {
		builder.WriteString(fmt.Sprintf(`"%v",` + "\n", scanner.Text()))
	}
	builder.WriteString("}")

	os.WriteFile("xx.txt", []byte(builder.String()), 0777)
	

}