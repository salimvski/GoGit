package main

import (
    "fmt"
    "os"
    "gogit/internal/commands"
)

func main() {

    if len(os.Args) < 2 {
        fmt.Fprintf(os.Stderr, "usage: gogit repo <list|stats> <username>\n")
        fmt.Fprintf(os.Stderr, "usage: gogit user <view|stats> <username>\n")
        os.Exit(1)
    }

    args := os.Args[1:]
    
    switch args[0] {
        case "user":
            commands.UserCmd(args[1:])
        case "repo":
            commands.RepoCmd(args[1:])
        default:
            fmt.Printf("unknown command: %s\n", os.Args[1])
    }
}