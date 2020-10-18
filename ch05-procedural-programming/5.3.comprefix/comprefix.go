package main

import (
	"bytes"
	"fmt"
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
	}

	for _, strings := range testData {
		fmt.Println(CommonPrefix(strings))
	}
}

func CommonPrefix(texts []string) string {
	if len(texts) == 0 {
		return ""
	}

	runeTexts := make([][]rune, len(texts))
	for i, text := range texts {
		runeTexts[i] = []rune(text)
	}

	var common bytes.Buffer
	for idx := 0; idx < len(runeTexts[0]); idx++ {
		for _, runeText := range runeTexts[1:] {
			if idx >= len(runeText) || runeText[idx] != runeTexts[0][idx] {
				return common.String()
			}
		}
		common.WriteRune(runeTexts[0][idx])
	}

	return common.String()
}
