// cmd/repo.go
package commands

import (
    "fmt"
    "os"
    "gogit/utils"
    "gogit/internal/api"
)

func RepoCmd(args []string) {

    if len(args) < 2 {
        fmt.Fprintf(os.Stderr, "usage: gogit repo <list|stats> <username>\n")
        os.Exit(1)
    }

    username := args[1]

    switch args[0] {
    case "list":
        user, err := api.GetUser(username)
        if err != nil {
            fmt.Fprintf(os.Stderr, "error: %v\n", err)
            os.Exit(1)
        }
        repos, err := api.GetUserRepos(user.ReposUrl)
		if err != nil {
            fmt.Fprintf(os.Stderr, "error: %v\n", err)
            os.Exit(1)
        }
		utils.PrintUserRepos(repos)
    case "stats":
        user, err := api.GetUser(username)
        if err != nil {
            fmt.Fprintf(os.Stderr, "error: %v\n", err)
            os.Exit(1)
        }
        repos, err := api.GetUserRepos(user.ReposUrl)
        if err != nil {
            fmt.Fprintf(os.Stderr, "error: %v\n", err)
            os.Exit(1)
        }
        utils.PrintUserStats(repos)
    default:
        fmt.Fprintf(os.Stderr, "unknown subcommand: %s\n", args[0])
        fmt.Fprintf(os.Stderr, "available: view, stats\n")
        os.Exit(1)
    }
}