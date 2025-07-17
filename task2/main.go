package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	letters := make(map[rune]int)

	fmt.Print("Enter your word: ")
	val, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error in parsing delimeter")
		return
	}
	val = strings.TrimSpace(val)

	for _, char := range val {
		lowerChar := unicode.ToLower(char)
		letters[lowerChar]++
	}
	for char, count := range letters {
		if count > 1 {
			fmt.Printf("%c repeats %d times\n", char, count)
		}
	}
}
