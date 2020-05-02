package gobasics

import "fmt"

func Mapsprog() {
	type empname struct {
		fname string
		lname string
	}

	emplist := make(map[string]empname)

	emplist["Chiranth"] = empname{"Chiranth", "C"}
	emplist["Rakesh"] = empname{"Rakesh", "V Shetty"}

	for mapkey, mapvalue := range emplist {
		fmt.Printf("Empkey : %v \n, Empdetails : $v \n ", mapkey, mapvalue)
	}
}
