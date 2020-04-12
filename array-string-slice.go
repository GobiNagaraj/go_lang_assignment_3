package main

import (
	"fmt"
)

func main () {
	fmt.Println("This program returns a given String in Substring as Slice")
	stringSlice()
}

func stringSlice() {
	str1 := [5]string{"Go", "is", "a", "Programming", "Language"}
	fmt.Printf("The String is : %v\n", str1)

	// return a string as slice
	sliceStr1 := str1[0:5]
	fmt.Println("The slice is : \n", sliceStr1)

	for i, v := range sliceStr1 {
		fmt.Println("Index %d: ", i)
		fmt.Println("Value %d: ", v)
		fmt.Println()
	}


}