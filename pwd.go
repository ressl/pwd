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
	test := "../"
	valid := true
	inode := 0
	dirs := ""
	pwd := ""

	for valid {
		inode, dirs = comparedir(test)
		if inode == 2 {
			pwd = fmt.Sprintf("/%v", pwd)
			valid = false
		} else {
			pwd = fmt.Sprintf("%v/%v", dirs, pwd)
			test += "../"
		}
	}
	fmt.Println(pwd)
}
