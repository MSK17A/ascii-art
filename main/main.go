package main

import (
	"asciiART"
	"fmt"
	"os"
)

func main() {
	//* handling invalid arguments
	if len(os.Args) < 2 {
		fmt.Print()
	} else if len(os.Args) >= 2 && len(os.Args) < 3 {
		//* handling null input text
		if os.Args[1] == "" {
			fmt.Print()
		} else {
			//* handling invalid charachters
			for _, v := range os.Args[1] {
				if v < ' ' || v > '~' {
					fmt.Println("Invalid charachters detected")
					return
				}
			}
			//* executing function
			asciiART.PrintART(os.Args[1], "standard")
		}
	} else if len(os.Args) >= 3 {
		//* executing function for multiple args
		if os.Args[1] == "" {
			return
		}
		asciiART.PrintART(os.Args[1], os.Args[2])
	}
}
