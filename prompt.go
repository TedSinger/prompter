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

type Prompt []*Part

func InitPrompt() Prompt {
    path, _ := os.Getwd()
    components := strings.Split(path, "/")
    prompt := make([]*Part, len(components))
    for i, _ := range prompt {
        part := Part{}
        part.Name = components[i]
        part.Abbreviation = components[i]
        part.NameStyle = make([]color.Attribute, 0, 4)
        prompt[i] = &part
    }
    return prompt
}

func (prompt Prompt) Format() string {
    ret := ""
    for idx, part := range prompt {
        ret += ApplyStyles(part.Abbreviation, part.NameStyle...)    
        is_last := idx == len(prompt) - 1
        is_abbrd := (part.Abbreviation != part.Name)
        is_tilde := (part.Abbreviation == "~" && part.Shadowed)
        if is_last {
            if is_abbrd && !is_tilde {
                ret += ApplyStyles("/", part.SlashStyle...)
            }
        } else if !part.Shadowed || is_tilde {
            ret += ApplyStyles("/", part.SlashStyle...)
        }
    }
    return ret
}
