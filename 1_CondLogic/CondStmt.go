package main

import (
	"log"
	"os"
	"text/template"
)

type college struct {
	Name string
	City string
}

type state struct {
	college
	Cond bool
}



func main() {

	st := state{
		college: college{
			Name: "Califoria state university",
			City: "Fresno",
		},
		Cond: false,
	}

	if st.Name == "Califoria state university"{
		st.Cond = true
	}else if st.Name == "Standford state university"{
		st.Cond = false
	}
	tpl, err := template.ParseFiles("college.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, st)
	if err != nil {
		log.Fatalln(err)
	}

}