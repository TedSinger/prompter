package main

import (
    "os"
    "fmt"
    "strings"
    "github.com/fatih/color"
)

func getStyles(p string) PathStyle {
    nComponents := strings.Count(p, "/") + 1
    styles := make(PathStyle, nComponents)
    for i := 0; i < nComponents; i += 1 {
        if i % 3 == 0 {
            styles[i] = []color.Attribute{color.FgRed}
        } else if i % 3 == 1 {
            styles[i] = []color.Attribute{color.FgHiBlue}
        } else if i % 3 == 2 {
            styles[i] = []color.Attribute{color.FgHiYellow, color.Underline}
        }
    }
    return styles
}

func getPrompt(p string) string {
    maxLen := getMaxPromptSize()
    abbrs := getAbbreviations(p, maxLen)
    styles := getStyles(p)
    return applyStyle(abbrs, styles)

}

func main() {
    path, _ := os.Getwd()
    fmt.Print(getPrompt(path))
}