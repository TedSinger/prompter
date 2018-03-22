package main

import (
    "os"
    "fmt"
    "strings"
    "path/filepath"
)


func getSubPaths(components []string) []string {
    subpaths := make([]string, len(components))
    current := ""
    for i, component := range components {
        current += string(filepath.Separator) + component
        subpaths[i] = current
    }
    return subpaths
}

func main() {
    path, _ := os.Getwd()
    components := strings.Split(path, string(filepath.Separator))
    // subpaths := getSubPaths(components)
    abbreviations := getAbbreviations(components, getMaxPromptSize())
    fmt.Print(strings.Join(abbreviations, string(filepath.Separator)))
}