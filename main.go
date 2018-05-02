package main

import (
    "os"
    "fmt"
    "strings"
    "github.com/fatih/color"
)

func getStyles(components []string) PathStyle {
    styles := make(PathStyle, len(components))
    last_fs_root := "/"
    current_fs_root := "/"
    mounts := getMounts()

    for i := 1; i < len(components); i += 1 {
        proper_path := strings.Join(components[:i+1], "/")
        resolved_path := proper_path
        styles[i] = make([]color.Attribute, 0, 4)

        // FIXME: use LSCOLORS instead of hardcoding
        if isLink(proper_path) {
            styles[i] = append(styles[i], color.FgHiCyan, color.Bold)
            resolved_path = resolvedPath(proper_path)
        } else {
            styles[i] = append(styles[i], color.FgBlue)

        }
        
        if isOpenWrite(proper_path) {
            styles[i] = append(styles[i], color.BgGreen)
        } else {
            styles[i] = append(styles[i], color.Bold)
        }

        current_fs_root = getPathRoot(resolved_path, mounts)
        if current_fs_root != last_fs_root {
            styles[i] = append(styles[i], color.Underline)
            last_fs_root = current_fs_root
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
    components := strings.Split(path, "/")
    fmt.Print(getPrompt(components))
}