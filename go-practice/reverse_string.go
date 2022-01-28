package main

import (
	"strings"
)

//reverse a string

func reverseString(s string) string {

	//make empty map
	r := []string{}
	//loop backwards on the string
	for i := len(s) - 1; i >= 0; i-- {
		//append each letter to reversed map
		r = append(r, string(s[i]))
	}
	//convert slice of strings to strins with Join()
	str := strings.Join(r, "")
	return str
}
