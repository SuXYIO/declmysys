package utils

import (
	"fmt"
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

		if len(str) != 1 {
			// too long
			fmt.Println("Input must be one character long")
			continue
		}

		switch str[0] {
		case 'Y', 'y':
			return true
		case 'N', 'n':
			return false
		default:
			continue
		}
	}
	// only to satisfy the compiler, theoretically unreachable
	return false
}
