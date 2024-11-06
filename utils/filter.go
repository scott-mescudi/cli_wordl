package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	var n = 2
	data := []byte{}
	file, err := os.Open("popular.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		word := scanner.Text()
		if len(word) == n {
			str := fmt.Sprintf("%v\n", word)
			for index := range str{
				data = append(data, str[index])
			}
		}
	}


	os.WriteFile(fmt.Sprintf("%v_letter_words.txt", n), data, 0700)
}