package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "path/filepath"
)

func TestAbbreviations(t *testing.T) {
    have := []string{"", "home", "qwghlm", "cryptonomicon"}
    want := []string{"", "h", "q", "crypton"}
    if len(string(filepath.Separator)) == 1 {
        assert.Equal(t, getAbbreviations(have, 12), want)    
    } else if len(string(filepath.Separator)) == 2 {
        assert.Equal(t, getAbbreviations(have, 15), want)    
    } 
}