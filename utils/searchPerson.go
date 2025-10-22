package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func SearchPerson(list []string, nama *string) (result []string) {
	defer func() {
		if r := recover(); r != nil {
			err := fmt.Sprintln(r)
			result = []string{err}
		}
	}()

	_, err := strconv.Atoi(*nama)

	if nama == nil {
		panic("insert can not empty")
	} else if err == nil {
		panic("insert can not a numerical")
	}

	for _, name := range list {
		nama0 := strings.ToLower(name)
		nama1 := strings.ToLower(*nama)
		if nama0 == nama1 {
			return []string{name}
		}
	}
	return []string{}
}
