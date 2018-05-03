package main
import (
    "regexp"
    "github.com/fatih/color"
)

type PathStyle [][]color.Attribute

func wrapWithSOHSTX(escapeCode string) string {
    return "\x01" + escapeCode + "\x02"
}

func explainZeroWidthEscapeCodesToGNUReadline(prompt string) string {
    escapeCodeFinder := regexp.MustCompile(`\x1b\[[0-9;]*m`)
    return escapeCodeFinder.ReplaceAllStringFunc(prompt, wrapWithSOHSTX)
}

func applyStyles(s string, styles ...color.Attribute) string {
    c := color.New(styles...)
    c.EnableColor()
    return c.Sprint(s)
}

func formatPath(components []string, abbrs []string, pathstyle PathStyle) string {
    prompt := ""
    startIndex := 0
    if startsWithUserHome(components) {
        prompt = applyStyles("~", color.FgYellow, color.Bold)
        if len(components) > 3 {
            prompt += "/"
        }
        startIndex = 3
    }

    for i := startIndex; i < len(abbrs); i++ {
        prompt += applyStyles(abbrs[i], pathstyle[i]...)
        if i != len(abbrs) - 1 {
            if abbrs[i] == components[i] {
                prompt += "/"
            } else {
                prompt += applyStyles("/", color.CrossedOut)
            }
        }
    }
    // FIXME: only do this for bash
    return explainZeroWidthEscapeCodesToGNUReadline(prompt)
}