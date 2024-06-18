package contacts

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func GetNextLineNumber(bookName string) (int, error) {
	file, err := os.Open(GetExecutablePath() + bookName + ".txt")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 1
	for scanner.Scan() {
		if scanner.Text() == "" {
			return lineNumber, nil
		}
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return lineNumber, nil
}

func CreateOrClearBook(name string, printInfo bool) {
	exPath := GetExecutablePath()
	filePath := exPath + name + ".txt"
	err := os.WriteFile(filePath, []byte(""), 0644)
	if err != nil {
		log.Fatal(err)
	} else {
		if printInfo {
			fmt.Println("Created new book at:", filePath)
		}
	}
}

func DoesBookExist(name string, printWarning bool) bool {
	exPath := GetExecutablePath()

	if _, err := os.Stat(exPath + name + ".txt"); errors.Is(err, os.ErrNotExist) {
		if printWarning {
			fmt.Println("Contact book", name, "does not exist in path", exPath)
		}
		return false
	}
	return true
}

func WriteLineToBook(name string, line string) {
	if !DoesBookExist(name, true) {
		return
	}

	exPath := GetExecutablePath()
	file, err := os.OpenFile(exPath+name+".txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	if _, err := file.WriteString(line + "\n"); err != nil {
		log.Fatal(err)
	}
}

func GetExecutablePath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath + "/"
}

func ReadAllContacts(book string) ([]Contact, error) {
	var result []Contact

	file, err := os.Open(GetExecutablePath() + book + ".txt")
	if err != nil {
		return result, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		if line == "" {
			break
		} else {
			entry, err := parseEntryFromLine(line)
			if err != nil {
				fmt.Println("Error creating Entry from line", strconv.Itoa(lineNumber))
			}
			result = append(result, entry)
		}
	}

	if err := scanner.Err(); err != nil {
		return result, err
	}

	return result, nil
}

func WriteAllContacts(book string, contacts []Contact) {
	for i := range contacts {
		writeContact(book, contacts[i])
	}
}

func writeContact(book string, contact Contact) {
	WriteLineToBook(book, "\""+strconv.Itoa(contact.Id)+"\";"+"\""+contact.FirstName+"\";\""+contact.LastName+"\";\""+contact.Email+"\";\""+contact.PhoneNumber+"\"")
}

func parseEntryFromLine(line string) (Contact, error) {
	parts := strings.Split(line, ";")

	if len(parts) != 5 {
		return Contact{}, fmt.Errorf("invalid line format")
	}

	id, err := strconv.Atoi(strings.Trim(parts[0], "\""))
	if err != nil {
		return Contact{}, err
	}

	firstName := strings.Trim(parts[1], "\"")
	lastName := strings.Trim(parts[2], "\"")
	email := strings.Trim(parts[3], "\"")
	phoneNumber := strings.Trim(parts[4], "\"")

	return Contact{
		Id:          id,
		FirstName:   firstName,
		LastName:    lastName,
		Email:       email,
		PhoneNumber: phoneNumber,
	}, nil
}

func RemoveEntryByIdAndChangeIdsAccordingly(entries []Contact, id int) []Contact {
	// Find index
	index := -1
	for i, entry := range entries {
		if entry.Id == id {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Printf("Entry with ID %d not found.\n", id)
		return entries
	}

	// Remove the entry
	entries = append(entries[:index], entries[index+1:]...)

	// Fix the ids of the following entries
	for i := index; i < len(entries); i++ {
		entries[i].Id -= 1
	}

	return entries
}
