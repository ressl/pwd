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
	fmt.Println(pwd())
}

func pwd() string {
	test := "../"
	valid := true
	inode := 0
	dirs := ""
	pwd := ""

	for valid {
		inode, dirs = comparedir(test)
		if inode == 2 {
			pwd = fmt.Sprintf("%v%v", dirs, pwd)
			valid = false
		} else {
			pwd = fmt.Sprintf("%v/%v", dirs, pwd)
			test += "../"
		}
	}
	fpwd := len(pwd) - 1
	return pwd[:fpwd]
}
