package main

/*
return true or false if phrase is a palindrome
*/

func is_plaindrome(str string) bool {

	//loop to 1+ the len/2 of the string
	for i := 0; i <= len(str)/2; i++ {
		//save the current char of string[i]
		curr := string(str[i])
		//check to see if the curr matches the same char at the same position in reverse
		if curr != string(str[len(str)-(1+i)]) {
			return false
		}
	}
	return true
}
