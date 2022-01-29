package main

import "fmt"

/*mutate list to add/remove from beginning or end

lst: list of values
command: either "remove", or "add"
location: either "begining" or "end"
value: when adding value to add

>>> list_manipulation(lst, 'add', 'beginning', 20)
        [20, 1, 2, 3]

>>> lst = [1, 2, 3]
>>> list_manipulation(lst, 'remove', 'end')
	[1, 2]

*/

func map_manipulation(lst []int, cmd string, loc string, n int) []int {
	newSlice := []int{}
	//If lst has no length then return  lst
	// if(len(lst) == 0 ){

	// }
	//if add, add n
	if cmd == "add" {
		if loc == "beginning" {
			//beginning, add new int to newLst, append the rest of lst
			newSlice = append(newSlice, n)
			fmt.Println("add n to beginning", newSlice)
			//Add slice to another slice
			newSlice = append(newSlice, lst...)
			fmt.Println("add lst to end", newSlice)
		} else {
			//end: append to slice and reassign to end
			newSlice = append(newSlice, lst...)
			fmt.Println("add to end", newSlice)
			newSlice = append(newSlice, n)
		}

	} else if cmd == "remove" {
		//if cmd is beginning, slice off beginning and reassign slice
		if loc == "beginning" {
			newSlice = lst[1:]
		} else {
			//if cmd is end , remove end and reassign slice
			newSlice = lst[:len(lst)-1]
		}

	}

	return newSlice

}
