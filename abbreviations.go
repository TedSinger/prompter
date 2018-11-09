package main

import (
	"os"
	"os/user"
	"strings"

	"github.com/fatih/color"
)

func getUserHome() []string {
	home := os.Getenv("HOME")
	if home != "" {
		return strings.Split(home, "/")
	}
	usr, err := user.Current()
	if err != nil {
		return []string{"non-matching sentinel value", "&\\!`\\\\$'"}
	} else {
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
		for i := range home {
			prompt[i].Shadowed = true
			prompt[i].Abbreviation = ""
		}
		prompt[len(home)-1].Abbreviation = "~"
		prompt[len(home)-1].NameStyle = []color.Attribute{color.FgYellow}
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

type TruncationTarget struct {
	maxSize  int
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

func getTruncationTarget(sizeCounts map[int]int, maxSize, charsToCut int) TruncationTarget {
	ret := TruncationTarget{
		maxSize,
		sizeCounts[maxSize],
		sizeCounts[maxSize-1],
	}
	for charsToCut > 0 && ret.maxSize > 1 {
		if ret.nMaxSize == 0 {
			ret.maxSize -= 1
			ret.nMaxSize = ret.nOneLess
			ret.nOneLess = sizeCounts[ret.maxSize-1]
		} else {
			ret.nMaxSize -= 1
			ret.nOneLess += 1
			charsToCut -= 1
		}
	}
	return ret
}

func GetTruncationTarget(prompt Prompt, charsToCut int) TruncationTarget {
	sizeCounts, maxSize := getSizeCounts(prompt)
	return getTruncationTarget(sizeCounts, maxSize, charsToCut)
}

func (tt *TruncationTarget) Decrement() {
	if tt.nOneLess > 0 {
		tt.nOneLess -= 1
	} else {
		tt.nMaxSize -= 1
	}
}

func (tt TruncationTarget) TargetSize() int {
	if tt.nOneLess > 0 {
		return tt.maxSize - 1
	} else {
		return tt.maxSize
	}
}

func SetAbbreviations(prompt Prompt, tt TruncationTarget) {
	for _, part := range prompt {
		if tt.nMaxSize > 0 && len(part.Abbreviation) == tt.maxSize {
			tt.nMaxSize -= 1
		} else if tt.nOneLess > 0 && len(part.Abbreviation) == tt.maxSize-1 {
			tt.nOneLess -= 1
		} else if len(part.Abbreviation) > tt.TargetSize() {
			part.Abbreviation = part.Abbreviation[:tt.TargetSize()]
			part.SlashStyle = []color.Attribute{color.CrossedOut}
			tt.Decrement()
		}
	}
}
