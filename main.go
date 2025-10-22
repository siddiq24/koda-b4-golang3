package main

import (
	"fmt"

	"github.com/siddiq24/golang3/utils"
)

func main() {
	list := []string{
		"Asep",
		"Bayu",
		"Caca",
		"Dadang",
		"Eko",
		"Fahrul",
		"Galih",
	}

	var search string
	fmt.Println("Insert Name to Search: ")
	_, err := fmt.Scan(&search)
	if err != nil {
		fmt.Print("your insert invalid")
		return
	}
	nama := utils.SearchPerson(list, search)
	fmt.Println(nama)
}