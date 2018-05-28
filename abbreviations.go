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
            prompt[i].Abbreviation = ""
        }
        prompt[len(home) - 1].Abbreviation = "~"
        prompt[len(home) - 1].NameStyle = []color.Attribute{color.FgYellow}
    }
}

func GetCharsToCut(prompt Prompt, maxSize int) int {
    totalChars := 0
    for _, part := range prompt {
        totalChars += len(part.Abbreviation)
        if !part.Shadowed {
            totalChars += 1
        }
    }
    var charsToCut int
    if totalChars > maxSize {
        charsToCut = totalChars - maxSize
    } else {
        charsToCut = 0
    }
    return charsToCut
}

type MaxSizeHolder struct {
    maxSize int
    nMaxSize int
    nOneLess int
}

func getSizeCounts(prompt Prompt) (map[int]int, int) {
    sizeCounts := make(map[int]int)
    maxSize := 0
    for _, part := range prompt {
        sizeCounts[len(part.Abbreviation)] += 1
        if len(part.Abbreviation) > maxSize {
            maxSize = len(part.Abbreviation)
        }
    }
    return sizeCounts, maxSize
}

func getMaxSizeHolder(sizeCounts map[int]int, maxSize, charsToCut int) MaxSizeHolder {
    ret := MaxSizeHolder{
        maxSize,
        sizeCounts[maxSize],
        sizeCounts[maxSize - 1],
    }
    for charsToCut > 0 && ret.maxSize > 1 {
        if ret.nMaxSize == 0 {
            ret.maxSize -= 1
            ret.nMaxSize = ret.nOneLess
            ret.nOneLess = sizeCounts[ret.maxSize - 1]
        } else {
            ret.nMaxSize -= 1
            ret.nOneLess += 1
            charsToCut -= 1
        }
    }
    return ret
}

func GetTargetMaxSize(prompt Prompt, charsToCut int) MaxSizeHolder {
    sizeCounts, maxSize := getSizeCounts(prompt)
    return getMaxSizeHolder(sizeCounts, maxSize, charsToCut)
}

func (msh *MaxSizeHolder) Decrement() {
    if msh.nOneLess > 0 {
        msh.nOneLess -= 1
    } else {
        msh.nMaxSize -= 1
    }
}

func (msh MaxSizeHolder) TargetSize() int {
    if msh.nOneLess > 0 {
        return msh.maxSize - 1
    } else {
        return msh.maxSize
    }
}

func SetAbbreviations(prompt Prompt, maxSizes MaxSizeHolder) {
    for _, part := range prompt {
        if maxSizes.nMaxSize > 0 && len(part.Abbreviation) == maxSizes.maxSize {
            maxSizes.nMaxSize -= 1
        } else if maxSizes.nOneLess > 0 && len(part.Abbreviation) == maxSizes.maxSize - 1 {
            maxSizes.nOneLess -= 1
        } else if len(part.Abbreviation) > maxSizes.TargetSize() {
            part.Abbreviation = part.Abbreviation[:maxSizes.TargetSize()]
            part.SlashStyle = []color.Attribute{color.CrossedOut}
            maxSizes.Decrement()
        }
    }
}