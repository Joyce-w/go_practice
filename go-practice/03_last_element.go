package main

import "errors"

// return last item in a list (Node if list is empty.)
//last_element([1,2,3]) //3
//last_element([]) //None

func last_element(a []int) (int, error) {

	if len(a) > 0 {
		return a[len(a)-1], nil
	}
	return 0, errors.New("None")

}

/* Alternative results, although cannot check for <nil>*/
// func last_element(a []int) (num int, str string) {

// 	if len(a) > 0 {
// 		num = a[len(a)-1]
// 	}
// 	str = "None"
// 	return num, str

// }
