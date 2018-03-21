package main


type Style int
type BgColor Style
type FgColor Style
type TextStyle Style


const (
    FgBlack FgColor = iota
    FgRed FgColor = iota
    FgGreen FgColor = iota
    FgYellow FgColor = iota
    FgBlue FgColor = iota
    FgMagenta FgColor = iota
    FgCyan FgColor = iota
    FgWhite FgColor = iota
    FgDefault FgColor = iota
)

const (
    BgBlack BgColor = iota
    BgRed BgColor = iota
    BgGreen BgColor = iota
    BgYellow BgColor = iota
    BgBlue BgColor = iota
    BgMagenta BgColor = iota
    BgCyan BgColor = iota
    BgWhite BgColor = iota
    BgDefault BgColor = iota
)

const (
    Reset TextStyle = iota
    Dim TextStyle = iota
    Standout TextStyle = iota
    Underscore TextStyle = iota
    Blink TextStyle = iota
    Reverse TextStyle = iota
    Hidden TextStyle = iota
)

func addStyle(str string, stys []Style) string {
    begin := ""
    for _, sty := range stys {
        begin += getEscapeCode(sty)
    }
    end := getEscapeCode(Style(Reset)) + getEscapeCode(Style(FgDefault)) + getEscapeCode(Style(BgDefault))
    return begin + str + end
}
