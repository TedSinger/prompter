package main

import (
    "fmt"
    "flag"
)

func main() {
    bash_hack := flag.Bool("bash-readline-hack", true, "Wraps escape codes in \\x01 and \\x02, so that GNU Readline understands that they have no width")
    default_color := flag.Int("default", 34, "color for normal directories (ANSI escape codes)")
    symlink_color := flag.Int("symlink", 36, "color for symlinks")
    open_write_color := flag.Int("open-write", 42, "color open write permissions")
    size_buffer := flag.Int("size-buffer", 10, "how much extra space to leave between the end of the prompt and the middle of the screen")
    flag.Parse()

    prompt := InitPrompt()
    ShadowHome(prompt)
    maxPromptSize := GetMaxPromptSize(*size_buffer)
    charsToCut := GetCharsToCut(prompt, maxPromptSize)
    maxSize := GetTargetMaxSize(prompt, charsToCut)
    SetAbbreviations(prompt, maxSize)
    StylePrompt(prompt, *default_color, *symlink_color, *open_write_color)
    if *bash_hack {
        // Bash can recognize escape codes in your PS1, but only if they are statically defined
        // When `prompter` prints them, Bash doesn't realize it has to explain them to GNU Readline
        // If your cursor jumps around when you scroll through your command history
        // then there's a bug here.
        fmt.Print(ExplainZeroWidthEscapeCodesToGNUReadline(prompt.Format()))
    } else {
        fmt.Print(prompt.Format())
    }    
}