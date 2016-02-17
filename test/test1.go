package main

import ("fmt"
	"io"
	"log"
	"os"
	"strings"

)

func main(){

name := "akram"

str := fmt.Sprint(`htmls data `+name+`end of the file data`)

	file,err :=os.Create("newfile.txt")
	if err != nil {
		log.Fatal("Error creating file",err)
}
defer file.Close()

io.Copy(file,strings.NewReader(str))



}
