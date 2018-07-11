package main

import (
	"fmt"
	"io/ioutil"
	"syscall"
)

func comparedir(path string) int {
	var cfile1 syscall.Stat_t
	var cfile2 syscall.Stat_t
	if path == "../" {
		syscall.Stat(".", &cfile1)
	} else if last := len(path) - 3; last >= 0 {
		syscall.Stat(path[:last], &cfile1)
	}
	dir, err := ioutil.ReadDir(path)
	check(err)

	tdir := ""
	for _, dirs := range dir {
		tdir = fmt.Sprintf("%v%v", path, dirs.Name())
		syscall.Stat(tdir, &cfile2)
		if cfile1.Ino == cfile2.Ino {
			return int(cfile2.Ino)
			break
		}
	}
	panic("error")
}
