package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	data := []byte{}
	file, err := os.Open("popular.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		word := scanner.Text()
		if len(word) == 5 {
			str := fmt.Sprintf("%v\n", word)
			for index := range str{
				data = append(data, str[index])
			}
		}
	}

	os.WriteFile("test.txt", data, 0700)
}