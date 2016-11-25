package main

import (
    "fmt"
    "os"
    "github.com/imattman/combo-chars-go/combo"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s {char-string}\n", os.Args[0])
        os.Exit(1)
    }
    
    combos := combo.FromRunes([]rune(os.Args[1]));
    for _, c := range combos {
        fmt.Printf("%s\n", c)
    }
}