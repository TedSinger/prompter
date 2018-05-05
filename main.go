package main

import (
    "fmt"
    "flag"
)

func main() {
    bash_hack := flag.Bool("bash-readline-hack", true, "Wraps escape codes in \\x01 and \\x02, so that GNU Readline understands that they have no width")
    default_color := flag.Int("default", 34, "color for normal directories (default blue)")
    symlink_color := flag.Int("symlink", 36, "color for symlinks (default cyan)")
    open_write_color := flag.Int("open-write", 42, "color open write permissions (default green background)")
    flag.Parse()

    prompt := InitPrompt()
    ShadowHome(prompt)
    charsToCut := GetCharsToCut(prompt)
    SetAbbreviations(prompt, charsToCut)
    StylePrompt(prompt, *default_color, *symlink_color, *open_write_color)
    if *bash_hack {
        fmt.Print(ExplainZeroWidthEscapeCodesToGNUReadline(prompt.Format()))
    } else {
        fmt.Print(prompt.Format())
    }
    
}