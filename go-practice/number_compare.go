package main

// report whether a>b, b>a, b==a
var (
	res string
)

func number_compare(a int, b int) string {

	if a > b {
		res = "first is greater"
	} else if b > a {
		res = "second is greater"
	} else {
		res = "Numbers are equal"
	}
	return res
}
