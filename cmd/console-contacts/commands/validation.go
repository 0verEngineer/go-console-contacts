package commands

import (
	"net/mail"
	"strings"
)

func nameRuneValid(r rune) bool {
	return !(r < 'A' || r > 'z')
}

func nameValid(name string) bool {
	if strings.IndexFunc(name, nameRuneValid) == -1 {
		return false
	}
	return true
}

func emailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func isStringValidForFile(s string) bool {
	if strings.Contains(s, "\"") {
		return false
	}
	return true
}
