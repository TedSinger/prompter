package main

import (
    "os"
    "fmt"
    "strings"
    "github.com/fatih/color"
    "path/filepath"
)

func getStyles(components []string) PathStyle {
    styles := make(PathStyle, len(components))
    for i := 0; i < len(components); i += 1 {
        if i % 3 == 0 {
            styles[i] = []color.Attribute{color.FgRed}
        } else if i % 3 == 1 {
            styles[i] = []color.Attribute{color.FgHiBlue}
        } else if i % 3 == 2 {
            styles[i] = []color.Attribute{color.FgHiYellow}
        }
        if i > 0 && isLink(strings.Join(components[:i+1], string(filepath.Separator))) {
            // FIXME: use LSCOLORS instead of hardcoding
            styles[i] = []color.Attribute{color.FgHiCyan, color.Bold}
        }
    }
    return styles
}

func getPrompt(components []string) string {
    maxLen := getMaxPromptSize()
    abbrs := getAbbreviations(components, maxLen)
    styles := getStyles(components)
    return applyStyle(abbrs, styles)
}

func main() {
    path, _ := os.Getwd()
    components := strings.Split(path, string(filepath.Separator))
    fmt.Print(getPrompt(components))
}