package main

/*return how many times a letter appears in a word*/
func singleLtrCount(ltr string, word string) int {
	count := 0
	for _, v := range word {

		if string(v) == ltr {
			count++
		}
	}

	return count
}
