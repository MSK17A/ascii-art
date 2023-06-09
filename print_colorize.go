package asciiART

import (
	"fmt"
)

func Print_Colorize(color string, input_str string) {
	var Colors = map[string]string{
		"black":   "\033[1;30m%s\033[0m",
		"red":     "\033[1;31m%s\033[0m",
		"green":   "\033[1;32m%s\033[0m",
		"yellow":  "\033[1;33m%s\033[0m",
		"blue":    "\033[1;34m%s\033[0m",
		"magenta": "\033[1;35m%s\033[0m",
		"cyan":    "\033[1;36m%s\033[0m",
		"white":   "\033[1;37m%s\033[0m",
	}
	fmt.Printf(Colors[color], input_str)
}
