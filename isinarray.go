package asciiART

/*
	inputs:
		num			number to be tested if it is in the array

	output:
		bool		(true or false)
*/

func isInArray(arr []int, num int) bool {

	for i := 0; i < len(arr); i++ {
		if num == arr[i] {
			return true
		}
	}
	return false
}
