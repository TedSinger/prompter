package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestAbbreviations(t *testing.T) {
    have := []string{"", "home", "qwghlm", "cryptonomicon"}
    want := []string{"", "h", "q", "crypton"}
    assert.Equal(t, getAbbreviations(have, 12), want)    
}