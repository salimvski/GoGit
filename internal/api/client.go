package api

import (
    "fmt"
    "io"
    "log"
    "net/http"
)


type Repo struct {
    Name        string `json:"name"`
    Description string `json:"description"`
    Stars       int    `json:"stargazers_count"`
    Forks       int    `json:"forks_count"`
}

func FetchUser(username string) ([]Repo, error) {
    // TODO: HTTP GET to GitHub API
	fetch_request := fmt.Sprintf("https://api.github.com/users/%s", username)

	req, err := http.NewRequest("GET", fetch_request, nil)
    if err != nil {
        log.Fatal(err)
    }

	req.Header.Add("Accept", "application/vnd.github+json")
    req.Header.Add("X-GitHub-Api-Version", "2022-11-28")

	client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    
    body, _ := io.ReadAll(resp.Body)
    fmt.Println(string(body))
    // TODO: Parse JSON response
	return nil, nil
}