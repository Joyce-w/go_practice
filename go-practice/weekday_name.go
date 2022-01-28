package main

// Return name of weekday
// weekday_name(1) //'Sunday'

func weekday_name(d int) string {

	// create a map with int key and string values
	m := map[int]string{
		1: "Sunday",
		2: "Monday",
		3: "Tuesday",
		4: "Wednesday",
		5: "Thursday",
		6: "Friday",
		7: "Saturday",
	}
	//return the key associated with the number
	return m[d]
}
