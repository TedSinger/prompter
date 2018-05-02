package main

import (
    "os"
    "os/exec"
    "fmt"
    "strings"
    "path/filepath"
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

func resolvedPath(p string) string {
    evaled, _ := filepath.EvalSymlinks(p)
    abs, _ := filepath.Abs(evaled)
    return abs
}

func isOpenWrite(p string) bool {
    lstat, err := os.Stat(p)
    if err != nil {
        // FIXME: have a debug mode, and only print with it
        fmt.Println(err.Error())
        return false
    } else {
        return (lstat.Mode() >> 1) & 1 == 1
    }   
}

func getMounts() []string {
    cmd := exec.Command("findmnt", "-l", "-o", "TARGET", "-n")
    bytes, _ := cmd.Output()
    text := string(bytes)
    return strings.Split(text, "\n")
}

func getPathRoot(p string, mounts []string) string {
    longest := "/"
    longest_size := 1
    for _, m := range mounts {
        if strings.HasPrefix(p, m) && len(m) > longest_size {
            longest = m
            longest_size = len(m)
        }
    }
    return longest
}