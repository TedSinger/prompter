package main
import (
    "os"
    "strings"
    "github.com/fatih/color"
)

type Part struct {
    Name string
    Abbreviation string
    NameStyle []color.Attribute
    SlashStyle []color.Attribute
    Shadowed bool
}

type Prompt []Part

func InitPrompt() Prompt {
    path, _ := os.Getwd()
    components := strings.Split(path, "/")
    prompt := make([]Part, len(components))
    for i, _ := range prompt {
        part := Part{}
        part.Name = components[i]
        part.NameStyle = make([]color.Attribute, 0, 4)
        prompt[i] = part
    }
    return prompt
}

func (prompt Prompt) Format() string {
    ret := ""
    start_idx := 0
    for start_idx < len(prompt) && prompt[start_idx].Shadowed {
        ret = ApplyStyles(prompt[start_idx].Abbreviation, prompt[start_idx].NameStyle...)
        start_idx += 1
    }
    for idx := start_idx; idx < len(prompt); idx ++ {
        if idx != 0 {
            ret += ApplyStyles("/", prompt[idx - 1].SlashStyle...)
        }
        
        ret += ApplyStyles(prompt[idx].Abbreviation, prompt[idx].NameStyle...)
    }
    last_part := prompt[len(prompt) - 1]
    if last_part.Abbreviation != last_part.Name && !last_part.Shadowed {
        ret += ApplyStyles("/", last_part.SlashStyle...)
    }
    return ret
}
