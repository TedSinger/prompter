package main

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func GetTermSize() (int, int) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	bytes, _ := cmd.Output()
	text := string(bytes)
	loc := strings.Index(text, " ")
	height, _ := strconv.Atoi(text[:loc])
	trimmed := strings.Trim(text[loc:], " \n\t\r")
	width, _ := strconv.Atoi(trimmed)
	return height, width
}

func GetMaxPromptSize(sizeBuffer int) int {
	_, w := GetTermSize()
	return w/2 - sizeBuffer
}
