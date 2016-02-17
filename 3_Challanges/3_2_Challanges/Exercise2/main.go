package main

import "fmt"

func main() {

half := func(n int) (int, bool) {
return n / 2, n%2 == 0
}

	var num int
	fmt.Println("Enter a Number")
	fmt.Scan(&num)
 	 h,even := half(num)
	fmt.Println("Entered Number/2 is :",h)
	fmt.Print("The Entered Number is :")
	if even {
		fmt.Println("Even")
	}else {
		fmt.Println("Not Even")
	}
}
