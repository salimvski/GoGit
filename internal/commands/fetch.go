package commands

import (
    "fmt"
    "gogit/internal/api"
)

func FetchUser(username string) {

    user, err := api.GetUser(username)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    fmt.Printf("╔══════════════════════════════════════╗\n")
    fmt.Printf("║        🐙 GITHUB PROFILE             ║\n")
    fmt.Printf("╠══════════════════════════════════════╣\n")
    fmt.Printf("║  • Username:  %-22s ║\n", user.Login)
    fmt.Printf("║  • Name:      %-22s ║\n", user.Name)
    fmt.Printf("║                                     ║\n")
    fmt.Printf("║  📦 Repositories: %-19d ║\n", user.PublicRepos)
    fmt.Printf("║  ⭐ Followers:    %-19d ║\n", user.Followers)
    fmt.Printf("║  🔄 Following:    %-19d ║\n", user.Following)
    fmt.Printf("╚══════════════════════════════════════╝\n")
}