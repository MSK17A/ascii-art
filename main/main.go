package main

import (
	"asciiART"
	"flag"
	"fmt"
	"os"
)

func main() {
	/*if len(os.Args) < 2 {
		fmt.Println("Usage: [OPTION] <apply options to this> <input string> [BANNER]")
		return
	} else if len(os.Args) >= 2 && len(os.Args) < 3 {
		if os.Args[1] == "" {
			fmt.Println("Usage: [OPTION] <apply options to this> <input string> [BANNER]")
			return
		} else {
			asciiART.PrintART("",os.Args[1], "standard", "")
		}
	} else if len(os.Args) >= 3 {
		//useColor := flag.Bool("color", false, "display colorized output")
		colorname := flag.String("color", "", "coloring")
		flag.Parse()
		if *colorname != "" {
			colorvalue := *colorname
			if len(os.Args) <= 3 {
				asciiART.PrintART("",os.Args[2], "standard", colorvalue)
			} else {
				asciiART.PrintART(os.Args[2],os.Args[3], "standard", colorvalue)
			}
		} else {
			asciiART.PrintART("",os.Args[1], os.Args[2], "")
		}
	}*/

	// Flags
	colorname := flag.String("color", "", "coloring")
	flag.Parse()

	if len(os.Args) == 2 && os.Args[1] != "" {
		asciiART.PrintART("",os.Args[1], "standard", "")
	} else if *colorname =="" && len(os.Args) == 3 && os.Args[2] != "" && os.Args[3] != ""{
		asciiART.PrintART("",os.Args[1], os.Args[2], "")
	} else if *colorname !="" && len(os.Args) == 3 {
		asciiART.PrintART("",os.Args[2], "standard", *colorname)
	} else if *colorname !="" && len(os.Args) == 4 {
		asciiART.PrintART(os.Args[2],os.Args[3], "standard", *colorname)
	} else if *colorname !="" && len(os.Args) == 5 {
		asciiART.PrintART(os.Args[2],os.Args[3], os.Args[4], *colorname)
	} else {
		fmt.Println("Usage: [OPTION] <apply options to this> <input string> [BANNER]")
	}
}