package core

import (
	"fmt"
	"strings"
	"time"
)

const maxAllowedGuesses = 10

func play(word string) {
	guessCount := 0
	guessWrongCount := 0
	guessed := map[string]bool{}
	guessList := []string{}
	guessProgress := initGuessProgress(word)

	gallows := newGallows()

	timeout := time.After(1 * time.Minute)
	ticker := time.NewTicker(500 * time.Millisecond)

	playing := true

	for playing {
		select {
		case <-timeout:
			phase := timesup

			fmt.Printf("\n%s\n", phase.desc)
			phase.run(gallows)
			gallows.render()

			fmt.Printf("Sorry, you lose. The word was %s\n", word)
			playing = false
		case <-ticker.C:
			fmt.Printf("\nHere are the letters you used: %s\n", strings.Join(guessList, ", "))
			fmt.Println(guessProgress)

			input := ""
			for input == "" {
				fmt.Print("What is your guess? ")
				fmt.Scanln(&input)

				input = sanitizeInput(input)
				_, alreadyGuessed := guessed[input]

				if len(input) > 0 && !isLetter(input) {
					fmt.Println("Only letters are allowed!")
					input = ""
				} else if alreadyGuessed {
					fmt.Println("You guessed that letter before!")
					input = ""
				} else {
					guessList = append(guessList, input)
					guessed[input] = true
					guessCount++
				}
			}

			if strings.Contains(word, input) {
				newProgress := strings.Split(guessProgress, "")
				for i, c := range word {
					if strings.ContainsRune(input, c) {
						newProgress[i] = string(c)
					}
				}

				guessProgress = strings.Join(newProgress, "")

				fmt.Printf("Letter %s is in the word\n", input)
				fmt.Println(guessProgress)

				if guessProgress == word {
					fmt.Println("You found the word!")
					playing = false
				}

				guessWord := ""
				fmt.Print("What is your guess for the word? ")
				fmt.Scanln(&guessWord)

				if guessWord == word {
					fmt.Printf("Correct! It took you %d guesses!\n", guessCount)
					playing = false
				}
			} else {
				phase := phases[guessWrongCount]

				fmt.Printf("\n%s\n", phase.desc)
				phase.run(gallows)
				gallows.render()

				guessWrongCount++

				fmt.Println("Sorry, that letter isn't in the word.")
			}

			if guessWrongCount == maxAllowedGuesses {
				fmt.Printf("Sorry, you lose. The word was %s\n", word)
				playing = false
			}
		}
	}
}

func initGuessProgress(word string) string {
	guessProgress := ""

	for range word {
		guessProgress += "-"
	}

	return guessProgress
}

func Start() {
	fmt.Println("HANGMAN")

	shuffleWords()
	wordIndex := 0

	var input string

	for {
		play(words[wordIndex])

		if wordIndex == len(words)-1 {
			fmt.Println("You did all the words!")
			break
		}

		fmt.Print("\nWant to play again? (Y or N) ")
		fmt.Scanln(&input)

		if sanitizeInput(input) != "Y" {
			break
		}

		wordIndex++
	}

	fmt.Println("It's been fun! Bye for now.")
}
