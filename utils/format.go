// utils/format.go
package utils

import (
    "fmt"
    "sort"
    "strings"
    "time"
	"gogit/internal/api"
)

func nullString(s string) string {
    if s == "" {
        return "-"
    }
    return s
}

func truncate(s string, n int) string {
    runes := []rune(s)
    if len(runes) <= n {
        return s
    }
    return string(runes[:n-1]) + "…"
}

func formatNumber(n int) string {
    if n < 1000 {
        return fmt.Sprintf("%d", n)
    }
    if n < 1_000_000 {
        return fmt.Sprintf("%.1fK", float64(n)/1000)
    }
    return fmt.Sprintf("%.1fM", float64(n)/1_000_000)
}

func formatDate(dateStr string) string {
    if dateStr == "" {
        return "-"
    }
    t, _ := time.Parse(time.RFC3339, dateStr)
    return t.Format("Jan 2006")
}

func PrintProfileCard(u api.User) {
    fmt.Printf("╔══════════════════════════════════════════╗\n")
    fmt.Printf("║           GITHUB PROFILE                ║\n")
    fmt.Printf("╠══════════════════════════════════════════╣\n")
    fmt.Printf("║  Photo: %s", u.AvatarURL)
    padding := 46 - len(u.AvatarURL)
    if padding > 0 {
        fmt.Print(strings.Repeat(" ", padding))
    }
    fmt.Printf("║\n")
    fmt.Printf("║                                          ║\n")
    fmt.Printf("║  Username      %-24s ║\n", "@"+u.Login)
    fmt.Printf("║  Name          %-24s ║\n", nullString(u.Name))
    fmt.Printf("║  Bio           %-24s ║\n", u.Bio)
    fmt.Printf("║  Location      %-24s ║\n", nullString(u.Location))
    fmt.Printf("║                                          ║\n")
    fmt.Printf("║  Public repos  %-24d ║\n", u.PublicRepos)
    fmt.Printf("║  Followers     %-24s ║\n", formatNumber(u.Followers))
    fmt.Printf("║  Following     %-24d ║\n", u.Following)
    fmt.Printf("║  Member since  %-24s ║\n", formatDate(u.CreatedAt))
    fmt.Printf("║                                          ║\n")
    fmt.Printf("║  %s", u.HTMLURL)
    if len(u.HTMLURL) < 42 {
        fmt.Print(strings.Repeat(" ", 42-len(u.HTMLURL)))
    }
    fmt.Printf(" ║\n")
    fmt.Printf("╚══════════════════════════════════════════╝\n")
}

func formatStars(n int) string {
    return formatNumber(n) + " ★"
}

func activityBadge(stars int) string {
    switch {
    case stars >= 1_000_000: return "LEGENDARY"
    case stars >= 250_000:   return "EXTREME"
    case stars >= 50_000:    return "VERY ACTIVE"
    case stars >= 10_000:    return "HIGHLY ACTIVE"
    case stars >= 1_000:     return "ACTIVE"
    case stars >= 100:       return "CASUAL"
    default:                 return "LOW ACTIVITY"
    }
}


func PrintUserStats(repos []api.Repo) {

    if len(repos) == 0 {
        fmt.Println("No repositories found.")
        return
    }

    totalStars := 0
    totalForks := 0
    mostStarredName := repos[0].Name
    mostStarredCount := repos[0].Stars

    langCount := make(map[string]int)

    for _, repo := range repos {
        totalStars += repo.Stars
        totalForks += repo.Forks

        if repo.Stars > mostStarredCount {
            mostStarredCount = repo.Stars
            mostStarredName = repo.Name
        }

        if repo.Language != "" {
            langCount[repo.Language]++
        }
    }

    // ─── Convert map → slice and sort by count ─────────────────────
    type langStat struct {
        name  string
        count int
    }
    var langList []langStat
    for lang, count := range langCount {
        langList = append(langList, langStat{name: lang, count: count})
    }
    sort.Slice(langList, func(i, j int) bool {
        return langList[i].count > langList[j].count
    })

    // ─── Print the beautiful stats card ───────────────────────────
    fmt.Printf("╔══════════════════════════════════════════╗\n")
    fmt.Printf("║          USER STATISTICS                ║\n")
    fmt.Printf("╠══════════════════════════════════════════╣\n")
    fmt.Printf("║                                          ║\n")
    fmt.Printf("║  Total stars earned   %-18s ║\n", formatStars(totalStars))
    fmt.Printf("║  Total repositories   %-18d ║\n", len(repos))
    fmt.Printf("║  Total forks earned   %-18s ║\n", formatNumber(totalForks))
    fmt.Printf("║                                          ║\n")
    fmt.Printf("║  Top languages                           ║\n")

    totalWithLang := len(repos)
    if totalWithLang == 0 {
        totalWithLang = 1
    }

    for i, l := range langList {
        if i >= 5 { break }
        percent := 100.0 * float64(l.count) / float64(totalWithLang)
        barLen := int(percent / 5)
        bar := strings.Repeat("█", barLen)

        fmt.Printf("║    • %-10s %5.1f%%  %s", l.name, percent, bar)
        fmt.Print(strings.Repeat(" ", 20-barLen))
        fmt.Printf(" ║\n")
    }

    fmt.Printf("║                                          ║\n")
    fmt.Printf("║  Most starred repo                       ║\n")
    fmt.Printf("║    → %-25s %6s ★   ║\n", truncate(mostStarredName, 25), formatNumber(mostStarredCount))
    fmt.Printf("║                                          ║\n")
    fmt.Printf("║  Activity level:      %-18s ║\n", activityBadge(totalStars))
    fmt.Printf("╚══════════════════════════════════════════╝\n")
}


func PrintUserRepos(repos []api.Repo) {
	fmt.Println()
	fmt.Printf("╔══════════════════════════════════════════════════════════════════════════════╗\n")
	fmt.Printf("║                            REPOSITORY LIST                                   ║\n")
	fmt.Printf("╚══════════════════════════════════════════════════════════════════════════════╝\n")
	fmt.Println()

	for i, repo := range repos {
		fmt.Printf("%d. %s\n", i+1, repo.Name)
		if repo.Description != "" {
			fmt.Printf("   %s\n", repo.Description)
		}
		fmt.Printf("   Language: %s | Stars: %d | Forks: %d\n", repo.Language, repo.Stars, repo.Forks)
		fmt.Println()
	}
}
