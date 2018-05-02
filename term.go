package main

import (
    "os"
    "os/exec"
    "strconv"
    "strings"
)

func getTermSize() (int, int) {
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

func getMaxPromptSize() int {
    _, w := getTermSize()
    return w / 2 - 10
}
