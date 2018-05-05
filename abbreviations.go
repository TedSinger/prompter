package main

import (
    "github.com/fatih/color"
)

func ShadowHome(prompt Prompt) {
    if StartsWithUserHome(prompt) {
        prompt[0].Shadowed = true
        prompt[1].Shadowed = true
        prompt[2].Shadowed = true
        prompt[2].Abbreviation = "~"
        prompt[2].NameStyle = []color.Attribute{color.FgYellow, color.Bold}
    }
}
func GetCharsToCut(prompt Prompt) int {
    maxLen := GetMaxPromptSize()
    totalChars := 0
    for i := 0; i < len(prompt); i++ {
        if prompt[i].Shadowed {
            totalChars += len(prompt[i].Abbreviation)
        } else {
            totalChars += 1 + len(prompt[i].Name)
        }
    }
    var charsToCut int
    if totalChars > maxLen {
        charsToCut = totalChars - maxLen
    } else {
        charsToCut = 0
    }
    return charsToCut
}
func SetAbbreviations(prompt Prompt, charsToCut int) {
    for i := 0; i < len(prompt); i++ {
        if prompt[i].Shadowed || prompt[i].Name == "" {
            continue
        } else if charsToCut >= len(prompt[i].Name) - 1 {
            prompt[i].Abbreviation = prompt[i].Name[:1]
            prompt[i].SlashStyle = []color.Attribute{color.CrossedOut}
            charsToCut -= len(prompt[i].Name) - 1
        } else if charsToCut > 0 {
            prompt[i].Abbreviation = prompt[i].Name[:len(prompt[i].Name) - charsToCut]
            prompt[i].SlashStyle = []color.Attribute{color.CrossedOut}
            charsToCut = 0
        } else {
            prompt[i].Abbreviation = prompt[i].Name
        }
    }
}
