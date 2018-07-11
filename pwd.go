package main

import (
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//comparedir("..")
	//test := "../../"
	//test := "../../../"
	test := "../../../"
	//for i := 2; i <= comparedir(test); test += "/.." {
	//	fmt.Println(i)
	//	fmt.Println(test)
	//test += "../"
	//comparedir(test)
	//	}
	fmt.Println(comparedir(test))
}
