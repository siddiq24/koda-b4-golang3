package utils

import (
	"strings"
)


func SearchPerson(list []string, nama *string) []string{
	for _,name := range list{
		nama0 := strings.ToLower(name)
		nama1 := strings.ToLower(*nama)
		if nama0 == nama1 {
			return []string{name}
		}
	}
	return []string{}
}