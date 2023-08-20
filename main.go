/*
Thematic Password Generator

The Thematic Password Generator is a customizable password generator that
allows users to create strong and unique passwords based on a predetermined
list of words related to various themes. This generator is intended for users
who want to create memorable passwords that are difficult to guess or crack.

The generator allows users to select a theme from a list including animals,
cars, Egyptian mythology, Greek/Roman mythology, Star Wars, technology, and
Tolkien. It then selects random keywords from the chosen theme and adds random
special characters and numbers to meet the desired length of the password.

This generator provides a user-friendly interface with prompts and menus that
guide the user through the password generation process. Additionally, it uses
the crypto/rand package to generate a cryptographically secure seed for the
math/rand package, ensuring the security of the generated passwords.

Author: Ryan Scrim
Date: 28/02/2023
Version: 1.0
*/

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var themes = map[string]string{
	"animals":               "animals.list",
	"cars":                  "cars.list",
	"egyptian_mythology":    "egyptian_mythology.list",
	"greek_roman_mythology": "greek_roman_mythology.list",
	"star_wars":             "star_wars.list",
	"technology":            "technology.list",
	"tolkein":               "tolkein.list",
}

func main() {
	fmt.Println("Welcome to the Password Generator!")
	for flowControl() {
		// Prompt user for theme choice
		fmt.Printf("Choose a theme: %s\n", strings.Join(getThemes(), ", "))
		theme := prompt("Enter a theme: ")

		// Read in list of keywords from file
		keywords, err := readKeywordsFromFile(themes[theme])
		if err != nil {
			handleError(err)
		}

		// Prompt user for number of passwords and password length
		numPasswords, pwLength := getNumPasswordsAndLength()

		// Generate specified number of passwords
		for i := 0; i < numPasswords; i++ {
			// Choose random keywords from list
			chosenKeywords := make([]string, 0)
			for j := 0; j < pwLength/4; j++ {
				chosenKeywords = append(chosenKeywords, strings.TrimSuffix(keywords[rand.Intn(len(keywords))], ","))
			}

			// Generate password
			password := ""
			for _, kw := range chosenKeywords {
				password += kw
				password += string(rand.Intn(10) + '0')      // add a random digit
				password += string(rand.Intn(26) + 'a')      // add a random lowercase letter
				password += string(rand.Intn(26) + 'A')      // add a random uppercase letter
				password += string("!@#$%^&*"[rand.Intn(8)]) // add a random special character
			}
			password = password[:pwLength] // trim to desired length

			// Print generated password
			fmt.Printf("Password %d: %s\n", i+1, password)
		}
	}
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func prompt(promptStr string) string {
	fmt.Print(promptStr)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func readKeywordsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	keywords := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		keywords = append(keywords, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return keywords, nil
}
func flowControl() bool {
	for {
		fmt.Println("\nWhat would you like to do?")
		fmt.Println("1. Generate a new password")
		fmt.Println("2. Quit")
		choice := prompt("Enter your choice: ")
		switch choice {
		case "1":
			return true
		case "2":
			clearScreen()
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func getNumPasswordsAndLength() (int, int) {
	// Read in number of passwords to generate
	numPasswordsStr := prompt("Enter the number of passwords to generate: ")
	numPasswords, err := strconv.Atoi(numPasswordsStr)
	if err != nil {
		handleError(fmt.Errorf("Invalid number of passwords entered: %v", numPasswordsStr))
	}

	// Read in desired length of passwords
	pwLengthStr := prompt("Enter the desired length of passwords: ")
	pwLength, err := strconv.Atoi(pwLengthStr)
	if err != nil {
		handleError(fmt.Errorf("Invalid password length entered: %v", pwLengthStr))
	}

	return numPasswords, pwLength
}

func getThemes() []string {
	themeChoices := make([]string, 0, len(themes))
	for theme := range themes {
		themeChoices = append(themeChoices, theme)
	}
	return themeChoices
}

func handleError(err error) {
	fmt.Printf("Error: %v\n", err)
	clearScreen()
}

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}
