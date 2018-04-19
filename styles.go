package main
import (
    "regexp"
    "github.com/fatih/color"
)

type PathStyle [][]color.Attribute

func explainZeroWidthEscapeCodesToGNUReadline(escapeCode string) string {
    return "\x01" + escapeCode + "\x02"
}

func applyStyle(abbrs []string, pathstyle PathStyle) string {
    prompt := ""
    var c *color.Color
    for i, abbr := range abbrs {
        c = color.New(pathstyle[i]...)
        c.EnableColor()
        if i == len(abbrs) - 1 {
            prompt += c.Sprint(abbr)
        } else {
            prompt += c.Sprint(abbr) + "/"
        }
    }
    escapeCodeFinder := regexp.MustCompile(`\x1b\[[0-9;]+m`)
    return escapeCodeFinder.ReplaceAllStringFunc(prompt, explainZeroWidthEscapeCodesToGNUReadline)
}