package main

import (
    "os"
    "fmt"
)

func isLink(p string) bool {
    lstat, err := os.Lstat(p)
    if err != nil {
        // FIXME: have a debug mode, and only print with it
        fmt.Println(err.Error())
        return false
    } else {
        return (lstat.Mode() & os.ModeSymlink) != 0
    }
}