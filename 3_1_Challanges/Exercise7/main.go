package main

import ( "fmt"
)

func main() {

	var sum int =0
	for index := 1; index < 1000; index++ {


		if index % 3==0 || index % 5 ==0{

			sum = sum+index

		}
	}
	fmt.Println("Sum is %d\n",sum)
}

