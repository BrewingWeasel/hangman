package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	file, err := os.Open("words.txt")
	checkErr(err)

	scanner := bufio.NewScanner(file)
	reader := bufio.NewReader(os.Stdin)

	var choices []string

	for scanner.Scan() {
		choices = append(choices, scanner.Text())
	}
	secretWord := choices[rand.Intn(len(choices))]

	curGuessed := make([]rune, len(secretWord))
	for i := range curGuessed {
		curGuessed[i] = '_'
	}

	numGuesses := 0

	for true {
		for _, char := range curGuessed {
			fmt.Print(string(char), " ")
		}
		fmt.Println("\nGuess a word or character: ")
		guessWord, err := reader.ReadString('\n')
		checkErr(err)
		guessWord = strings.Trim(guessWord, "\n")

		if len(guessWord) != len(secretWord) && len(guessWord) != 1 {
			fmt.Println("Wrong number of characters!")
			continue
		}

		if guessWord == secretWord {
			break
		} else if len(guessWord) == 1 {
			guessChar := []rune(guessWord)[0]
			for i, char := range secretWord {
				if char == guessChar {
					curGuessed[i] = char
				}
			}
		}

		if string(curGuessed) == secretWord {
			break
		}
		numGuesses++
	}
	fmt.Println("You win! It took you ", numGuesses, " guesses.")
}

func init() {
	rand.Seed(time.Now().Unix())
}
