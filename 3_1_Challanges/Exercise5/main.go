package main

import ( "fmt"
)

func main() {

	fmt.Println ("Printing All Even Numbers")
	for index := 0; index <= 100; index++ {

		if index % 2==0 {

			fmt.Println(index)
		}
	}
}

