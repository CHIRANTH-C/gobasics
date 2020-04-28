package gobasics

import "fmt"

func ArrayForLoop() {
	var names [6]string
	names = [6]string{"Jim", "Frank", "Troy", "Marry", "Jenny", "Anny"}

	fmt.Println("*******************")
	fmt.Println("Pass By Value")
	fmt.Println("*******************")

	for i, v := range names {
		fmt.Println(i, v)
	}
	fmt.Println("*******************")
	fmt.Println("Pass By Referance")
	fmt.Println("*******************")
	for i := range names {
		fmt.Println(names[i])
	}
}
