package main

import ( "fmt"
)

func main(){

	var fNumber int
	var sNumber int
	fmt.Println("Enter First Number")
	fmt.Scan(&fNumber)
	fmt.Println (" Please Enter Seconnd Number")
	fmt.Scan(&sNumber)
	fmt.Println("Remainder of Big Number / Small Number is :")
	if fNumber > sNumber {

		fmt.Println(fNumber % sNumber)
	}else {
		fmt.Println (sNumber % fNumber)
	}
	}

