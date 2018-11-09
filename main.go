package main

import (
	"flag"
	"fmt"
)

func main() {
	bashHack := flag.Bool("bash-readline-hack", true, "Wraps escape codes in \\x01 and \\x02, so that GNU Readline understands that they have no width")
	defaultColor := flag.Int("default", 34, "color for normal directories (ANSI escape codes)")
	symlinkColor := flag.Int("symlink", 36, "color for symlinks")
	openWriteColor := flag.Int("open-write", 42, "color open write permissions")
	sizeBuffer := flag.Int("size-buffer", 10, "how much extra space to leave between the end of the prompt and the middle of the screen")
	flag.Parse()

	prompt := InitPrompt()
	ShadowHome(prompt)
	maxPromptSize := GetMaxPromptSize(*sizeBuffer)
	charsToCut := GetCharsToCut(prompt, maxPromptSize)
	truncationTarget := GetTruncationTarget(prompt, charsToCut)
	SetAbbreviations(prompt, truncationTarget)
	StylePrompt(prompt, *defaultColor, *symlinkColor, *openWriteColor)
	if *bashHack {
		// Bash can recognize escape codes in your PS1, but only if they are statically defined
		// When `prompter` prints them, Bash doesn't realize it has to explain them to GNU Readline
		// If your cursor jumps around when you scroll through your command history
		// then there's a bug here.
		fmt.Print(ExplainZeroWidthEscapeCodesToGNUReadline(prompt.Format()))
	} else {
		fmt.Print(prompt.Format())
	}
}
