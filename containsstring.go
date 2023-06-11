package asciiART

import "strings"

/*
	This function will test if the letter to be printed is to be colored or not
	inputs:
		sub_str:		letters to be colored
		input_str:	the letter to be tested

	output:
		index			index of the first letter
*/

func ContainsString(sub_str string, input_str string) (int, int) {
	index := strings.Index(input_str, sub_str)

	return index, len(sub_str)
}
