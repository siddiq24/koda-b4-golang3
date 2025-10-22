package utils

import (
	"fmt"
	"strings"
	"unicode"
)

func SearchPerson(list []string, nama *string) (result []string) {
	defer func() {
		if r := recover(); r != nil {
			err := fmt.Sprintln(r)
			result = []string{err}
		}
	}()

	if nama == nil || strings.TrimSpace(*nama) == "" {
		panic("insert can not be empty")
	}

	for _, r := range *nama {
		if unicode.IsDigit(r) {
			panic("input can not include a number")
		}
	}

	input := strings.ToLower(*nama)
	for _, name := range list {
		if strings.ToLower(name) == input {
			return []string{name}
		}
	}

	return []string{}
}
