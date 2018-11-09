package main

import (
	"regexp"
)

func wrapWithSOHSTX(escapeCode string) string {
	return "\x01" + escapeCode + "\x02"
}

func ExplainZeroWidthEscapeCodesToGNUReadline(prompt string) string {
	escapeCodeFinder := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	return escapeCodeFinder.ReplaceAllStringFunc(prompt, wrapWithSOHSTX)
}
