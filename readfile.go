package gobasics

import (
	"fmt"
	"io/ioutil"
)

// This is ReadThisFile method
func ReadThisFile() {
	data, err := ioutil.ReadFile("myfile.txt")
	if err != nil {
		fmt.Println("Could not read file with error : ", err)
		return
	}
	fmt.Println("content of the file is : ", string(data))
}
