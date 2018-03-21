package main

import (
    "path/filepath"
    "strings"
)

func toAbbreviations(p string, maxLen int) []string {
    parts := strings.Split(p, string(filepath.Separator))
    var charsToCut int
    if len(p) > maxLen {
        charsToCut = len(p) - maxLen
    } else {
        charsToCut = 0
    }
    clips := make([]int, len(parts))
    current := ""
    for i, part := range parts {
        current += string(filepath.Separator) + part
        if i != 0 {
            if charsToCut >= len(part) - 1 {
                clips[i] = 1
                charsToCut -= len(part) - 1
            } else if charsToCut > 0 {
                clips[i] = len(part) - charsToCut
                charsToCut = 0
            } else {
                clips[i] = len(part)
            }
        }
    }
    ret := make([]string, len(parts))
    for i, part := range parts {
        ret[i] = part[:clips[i]]
    }
    return ret
}
