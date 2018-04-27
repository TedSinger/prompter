package main

import "path/filepath"

func getAbbreviations(components []string, maxLen int) []string {
    totalChars := 0
    for i, component := range components {
        if i != 0 {
            totalChars += len(string(filepath.Separator)) + len(component)    
        }
    }
    var charsToCut int
    if totalChars > maxLen {
        charsToCut = totalChars - maxLen
    } else {
        charsToCut = 0
    }
    abbrs := make([]string, len(components))
    for i, component := range components {
        if i != 0 {
            if charsToCut >= len(component) - 1 {
                abbrs[i] = component[:1]
                charsToCut -= len(component) - 1
            } else if charsToCut > 0 {
                abbrs[i] = component[:len(component) - charsToCut]
                charsToCut = 0
            } else {
                abbrs[i] = component
            }
        }
    }
    return abbrs
}
