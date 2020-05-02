package gobasics

import "fmt"

func Slicespro() {
	studentnames := []string{"Chiranth", "Hemanth", "Rakesh"}
	fmt.Println{"*******************"}
	for index, value := range studentnames {
		fmt.Printf("Rol No : %v - Name : %v", index, value)
	}
}
