package main

// type counter struct {
// 	letter string
// 	count  int
// }

func letterCount(w string) map[string]int {
	c := map[string]int{}

	//loop through the word returning char code
	for _, v := range w {
		//if counter has the letter from char code, increment count
		_, ok := c[string(v)]
		//check to see if key is in map, if not false (making statement true), add letter and count

		if ok == false {
			// fmt.Println("this is false")
			c[string(v)]++
		} else {
			c[string(v)]++
		}

	}
	return c
}
