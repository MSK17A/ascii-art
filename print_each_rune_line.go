package asciiART

import "fmt"

func Print_Each_Rune_Line(letters_to_be_colored string, str string, fontname string, color string) {

	string_first_char_idx, sub_len := ContainsString(letters_to_be_colored, str)
	coloring_enable := false
	end_str := 0

	fmt.Println(string_first_char_idx)
	// if there is parts to be colored
	if letters_to_be_colored != "" {

		//* Iterate through eight lines
		for i := 0; i < 8; i++ {
			//* Iterate through each character in the input string
			str_len := len(str)
			for idx := 0; idx < str_len; idx++ {
				char := str[idx]
				//* Check for the backslash (since this is taken from os arguments you have to read it like this '\\') and then read the letter after it
				if char == '\\' {
					if idx < len(str)-1 {
						//* Apply tab if 't' appears
						if str[idx+1] == 't' {
							fmt.Print("\t")
							idx++
						} else {
							if isInArray(string_first_char_idx, idx) || (coloring_enable && idx < end_str) {
								if !coloring_enable {
									end_str = idx + sub_len
									coloring_enable = true
								}
								// Start printing the colored letler ART
								PrintFileLine(MapART(rune(char))+i, MapFont(fontname), color)

							} else {
								// Start printing the letler ART in default color
								PrintFileLine(MapART(rune(char))+i, MapFont(fontname), "")
							}
						}
					}

				} else {
					if isInArray(string_first_char_idx, idx) || (coloring_enable && idx < end_str) {
						if !coloring_enable {
							end_str = idx + sub_len
							coloring_enable = true
						}
						// Start printing the colored letler ART
						PrintFileLine(MapART(rune(char))+i, MapFont(fontname), color)

					} else {
						// Start printing the letler ART in default color
						PrintFileLine(MapART(rune(char))+i, MapFont(fontname), "")
					}
				}
			}
			coloring_enable = false
			fmt.Print("\n") //* prints newline to start printing the rest of the art
		}
	} else {
		//* Iterate through eight lines
		for i := 0; i < 8; i++ {
			//* Iterate through each character in the input string
			str_len := len(str)
			for idx := 0; idx < str_len; idx++ {
				char := str[idx]
				//* Check for the backslash (since this is taken from os arguments you have to read it like this '\\') and then read the letter after it
				if char == '\\' {
					if idx < len(str)-1 {
						//* Apply tab if 't' appears
						if str[idx+1] == 't' {
							fmt.Print("\t")
							idx++
						} else {
							// Start printing the colored letler ART
							PrintFileLine(MapART(rune(char))+i, MapFont(fontname), color)
						}
					}
				} else {
					// Start printing the colored letler ART
					PrintFileLine(MapART(rune(char))+i, MapFont(fontname), color)
				}
			}
			fmt.Print("\n") //* prints newline to start printing the rest of the art
		}
	}
}
