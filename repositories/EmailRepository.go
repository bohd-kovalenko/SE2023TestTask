package repositories

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
)

const EmailBaseFileName = "emailBase.txt"
const ValidEmailPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

var ErrInvalidEmail = errors.New("invalid email format")

func SaveEmailToFile(email string) error {
	err := validateEmail(email)
	if err != nil {
		fmt.Println("Email is not valid ", err)
		return err
	}
	file, err := os.OpenFile(EmailBaseFileName, os.O_APPEND, 2)
	if err != nil {
		fmt.Println("Error when opening the file", err)
		return err
	}
	defer file.Close()
	_, err = file.Write([]byte(email + "\n"))
	if err != nil {
		fmt.Println("Error when writing to the file")
		return err
	}
	return nil
}

func FindEmailInFile(email string) (bool, error) {
	file, err := os.OpenFile(EmailBaseFileName, os.O_APPEND, 4)
	if err != nil {
		fmt.Println("Error, when opening file ", err)
		return false, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == email {
			return true, nil
		}
	}
	return false, nil
}

func FindAllEmails() ([]string, error) {
	var result []string
	file, err := os.OpenFile(EmailBaseFileName, os.O_APPEND, 4)
	if err != nil {
		fmt.Println("Error, when opening file", err)
		return result, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result, nil
}
func validateEmail(email string) error {
	isMatch, err := regexp.Match(ValidEmailPattern, []byte(email))
	if err != nil {
		fmt.Println("Error in pattern matching ", err)
		return err
	}
	if !isMatch {
		fmt.Println("Email does not match the pattern ")
		return ErrInvalidEmail
	}
	return nil
}
