package main

import "fmt"

func max(arr []int) int {
	var largest int
	for _, v := range arr {
		if v > largest {
			largest = v
		}
	}
	return largest
}

func main() {

	var length int
	fmt.Println("Enter the length of an array")
	fmt.Scan(&length)
	 var a = make([]int,length)
	var temp int
	fmt.Println("Enter ",length," values of the array")
	for i :=0;i < length; i++ {
		fmt.Println("Enter Value no :",i+1)
		fmt.Scan(&temp)
		a[i] = temp
	}




	greatest := max (a)
	//greatest := max(4, 7, 9, 123, 543, 23, 435, 53, 125)
	fmt.Println("The greatest of the entered numbers is:")
	fmt.Println(greatest)
}
