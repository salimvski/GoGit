// cmd/repo.go
package commands

import (
    "fmt"
    "os"
    "gogit/utils"
    "gogit/internal/api"
    "gogit/internal/models"
    "sync"
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

        var allRepos []models.RepoWithContributors

        var wg sync.WaitGroup

        inputRepoChan := make(chan models.Repo)
        outputRepoChan := make(chan models.RepoWithContributors)

        go func() {
            for _, repo := range repos {
                inputRepoChan <- repo
            }
            close(inputRepoChan)
        }()

        for i:= 0; i < 5; i++ {
            wg.Add(1)

            go func() {
                defer wg.Done()

                for repo := range inputRepoChan {
                    contributors, err := api.GetRepoContributors(repo.ContributorsURL)

                    if err != nil {
                        fmt.Printf("Error fetching contributors: %v\n", err)
                        continue
                    }

                    outputRepoChan <- models.RepoWithContributors{
                        RepoName:     repo.Name,
                        Contributors: contributors,
                    }

                }
            }()
        }

        go func() {
            wg.Wait()
            close(outputRepoChan)
        }()

        for result := range outputRepoChan {
            allRepos = append(allRepos, result)
        }

        utils.PrintAllReposWithContributors(allRepos)
    

    default:
        fmt.Fprintf(os.Stderr, "unknown subcommand: %s\n", args[0])
        fmt.Fprintf(os.Stderr, "available: list, stats\n")
        os.Exit(1)
    }
}