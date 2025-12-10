package api

import (
	"encoding/json"
    "fmt"
    "io"
    "net/http"
)

type User struct {
    Login       string `json:"login"`
    Name        string `json:"name"`
    PublicRepos int    `json:"public_repos"`
	ReposUrl 	string `json:"repos_url"`
    Followers   int    `json:"followers"`
    Following   int    `json:"following"`
}

func GetUser(username string) (User, error) {

	url := fmt.Sprintf("https://api.github.com/users/%s", username)

	req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return User{}, fmt.Errorf("creating request: %w", err)
    }

	req.Header.Add("Accept", "application/vnd.github+json")
    req.Header.Add("X-GitHub-Api-Version", "2022-11-28")

	client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return User{}, fmt.Errorf("making request: %w", err)
    }
    defer resp.Body.Close()

	if resp.StatusCode != 200 {
        return User{}, fmt.Errorf("API returned status: %d", resp.StatusCode)
    }
    
    body, err := io.ReadAll(resp.Body)
	if err != nil {
        return User{}, fmt.Errorf("reading response: %w", err)
    }

	var user User

	if err := json.Unmarshal(body, &user); err != nil {
        return User{}, fmt.Errorf("parsing JSON: %w", err)
    }

	return user, nil
}

type Repo struct {
    Name        string `json:"name"`
    Description string `json:"description"`
	Language 	string `json:"language"`
    Stars       int    `json:"stargazers_count"`
    Forks       int    `json:"forks_count"`
}

func GetUserRepos(repos_url string) ([]Repo, error) {

	url := repos_url

	req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, fmt.Errorf("creating request: %w", err)
    }

	req.Header.Add("Accept", "application/vnd.github+json")
    req.Header.Add("X-GitHub-Api-Version", "2022-11-28")

	client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("making request: %w", err)
    }
    defer resp.Body.Close()

	if resp.StatusCode != 200 {
        return nil, fmt.Errorf("API returned status: %d", resp.StatusCode)
    }
    
    body, err := io.ReadAll(resp.Body)
	if err != nil {
        return nil, fmt.Errorf("reading response: %w", err)
    }

	var repos []Repo

	if err := json.Unmarshal(body, &repos); err != nil {
        return nil, fmt.Errorf("parsing JSON: %w", err)
    }

	// for _, repo := range repos {
    //     fmt.Printf("Name: %s, Description: %s, Language: %s\n", repo.Name, repo.Description, repo.Language)
    // }

	return repos, nil
}