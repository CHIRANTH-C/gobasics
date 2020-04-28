package gobasics

import "fmt"

func ArrayForLoop() {
	var names [6]string
	names = [6]string{"Jim", "Frank", "Troy", "Marry", "Jenny", "Anny"}
	for i, v := range names {
		fmt.Println(i, v)
	}
}
