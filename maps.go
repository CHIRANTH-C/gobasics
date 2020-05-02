package gobasics

import "fmt"

func Mapsprog() {
	type empname struct {
		fname string
		lname string
	}

	fmt.Println()

	fmt.Println("*****************************")

	emplist := make(map[string]empname)

	emplist["Chiranth"] = empname{"Chiranth", "C"}
	emplist["Rakesh"] = empname{"Rakesh", "V Shetty"}

	for mapkey, mapvalue := range emplist {
		fmt.Printf("Empkey : %s \n", mapkey)
		fmt.Println(mapvalue)
	}
	fmt.Println()
	fmt.Println("*****************************")
}
