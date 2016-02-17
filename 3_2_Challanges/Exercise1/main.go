package main

import ( "fmt")

func main(){
	var n int
	fmt.Println ("Enter a Number")
	fmt.Scan(&n)
	h, even := half(n)
	fmt.Println("Entered Number/2 is :",h)
	fmt.Print("The Entered Number is :")
	if even {
		fmt.Println("Even")
	}else {
		fmt.Println("Not Even")
	}
}
func half(v int) (int, bool) {
	return v / 2, v%2 == 0
}