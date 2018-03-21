package main

import (
    "os"
    "fmt"
    "strings"
)


func toPrompt(p string) string {
    maxLen := getMaxPromptSize()
    return strings.Join(toAbbreviations(p, maxLen), "/")
}

func main() {
    path, _ := os.Getwd()
    fmt.Print(toPrompt(path))
}