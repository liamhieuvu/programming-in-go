package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	testData := [][]string{
		{"/home/user/goeg", "/home/user/goeg/prefix",
			"/home/user/goeg/prefix/extra"},
		{"/home/user/goeg", "/home/user/goeg/prefix",
			"/home/user/prefix/extra"},
		{"/pecan/π/goeg", "/pecan/π/goeg/prefix",
			"/pecan/π/prefix/extra"},
		{"/pecan/π/circle", "/pecan/π/circle/prefix",
			"/pecan/π/circle/prefix/extra"},
		{"/home/user/goeg", "/home/users/goeg",
			"/home/userspace/goeg"},
		{"/home/user/goeg", "/tmp/user", "/var/log"},
		{"/home/mark/goeg", "/home/user/goeg"},
		{"home/user/goeg", "/tmp/user", "/var/log"},
		{"home/user/goeg", "home/tmp/user", "home/var/log"},
	}

	for _, strings := range testData {
		fmt.Println(CommonPathPrefix(strings))
	}
}

func CommonPathPrefix(paths []string) string {
	const separator = string(filepath.Separator)
	dirs := make([][]string, len(paths))
	for i, p := range paths {
		dirs[i] = strings.Split(p, separator)
	}

	if len(dirs) == 0 || len(dirs[0]) == 0 {
		return ""
	}

	var col int
END:
	for col = 0; col < len(dirs[0]); col++ {
		for _, d := range dirs[1:] {
			if col >= len(d) || d[col] != dirs[0][col] {
				break END
			}
		}
	}
	if dirs[0][0] == "" {
		dirs[0][0] = separator
	}
	return filepath.Join(dirs[0][:col]...)
}
