package main

func getTermSize() int {
    return 120
}

func getMaxPromptSize() int {
    return getTermSize() / 2 - 20
}

func getEscapeCode(s Style) string {
    return ""
}