package main

import (
	"log"
	"os"
	"text/template"
)

type college struct {
	Name string
	city string
}

type state struct {
	college
	cond bool
}



func main() {

	st := state{
		college: college{
			Name: "Califoria state university",
			city: "Fresno",
		},
		cond: false,
	}

	if st.Name == "Califoria state university"{
		st.cond = true
	}else if st.Name == "Standford state university"{
		st.cond = false
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