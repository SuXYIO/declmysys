package utils

import (
	"fmt"
	"strings"
)

// Really should pass a option type for defaults and valid chars and stuff, but i dont think i need those functionalities

// AskYN asks yes or no until got valid
func AskYN(msg string) bool {
	// true for yes, false for no
	/*
		usage example:
			ans := AskYNQ("delete file")
			if ans == true {
				// delete
			} else {
				// don't delete
			}
	*/
	for true {
		var str string
		fmt.Print(msg + " [y/n]: ")
		fmt.Scanln(&str)

		switch strings.ToLower(str) {
		case "y", "yes":
			return true
		case "n", "no":
			return false
		default:
			fmt.Println(`Must be one of "y" "yes" "n" "no" (case-insensitive)`)
			continue
		}
	}
	// only to satisfy the compiler, theoretically unreachable
	return false
}
