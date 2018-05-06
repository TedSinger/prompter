package main

import (
    "github.com/fatih/color"
    "os/user"
)

func getUserHome() []string {
    usr, err := user.Current()
    if err != nil {
        return []string{"non-matching sentinel value", "&\\!`\\\\$'"}
    } else {
        // FIXME: it is not always true that a user's home directory is /home/username
        return []string{"", "home", usr.Username}
    }
}

func promptStartsWith(prompt Prompt, prefix []string) bool {
    if len(prefix) > len(prompt) {
        return false
    }
    for i, name := range prefix {
        if name != prompt[i].Name {
            return false
        }
    }
    return true
}

func ShadowHome(prompt Prompt) {
    home := getUserHome()
    if promptStartsWith(prompt, home) {
        for i, _ := range home {
            prompt[i].Shadowed = true
        }
        prompt[len(home) - 1].Abbreviation = "~"
        prompt[len(home) - 1].NameStyle = []color.Attribute{color.FgYellow}
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
