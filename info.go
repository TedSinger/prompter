package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func IsLink(p string) bool {
	lstat, err := os.Lstat(p)
	if err != nil {
		// FIXME: have a debug mode, and only print with it
		fmt.Println(err.Error())
		return false
	} else {
		return (lstat.Mode() & os.ModeSymlink) != 0
	}
}

func ResolvedPath(p string) string {
	evaled, _ := filepath.EvalSymlinks(p)
	abs, _ := filepath.Abs(evaled)
	return abs
}

func IsOpenWrite(p string) bool {
	lstat, err := os.Stat(p)
	if err != nil {
		// FIXME: have a debug mode, and only print with it
		fmt.Println(err.Error())
		return false
	} else {
		return (lstat.Mode()>>1)&1 == 1
	}
}

func GetMounts() []string {
	cmd := exec.Command("findmnt", "-l", "-o", "TARGET", "-n")
	bytes, _ := cmd.Output()
	text := string(bytes)
	return strings.Split(text, "\n")
}

func GetPathRoot(p string, mounts []string) string {
	longest := "/"
	longestSize := 1
	for _, m := range mounts {
		if strings.HasPrefix(p, m) && len(m) > longestSize {
			longest = m
			longestSize = len(m)
		}
	}
	return longest
}
