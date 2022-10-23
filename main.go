package main

import "main/my_hangman/src"

func main() {
	hangman := src.NewHangManData()
	for hangman.Attempts > 0 {
		hangman.AskForALetter()
	}

	if hangman.Attempts == 0 {
		hangman.PrintLose()
	}

}
