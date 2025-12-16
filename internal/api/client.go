package api

import (
	"encoding/json"
    "fmt"
    "github.com/joho/godotenv"
    "io"
    "net/http"
    "os"
    "gogit/internal/models"
)

var githubToken string
var defaultHeaders map[string]string

func init() {
    err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file:", err)
    }
    
    // Now get the token
    githubToken = os.Getenv("GITHUB_TOKEN")
    
    // Now create headers with the token
    defaultHeaders = map[string]string{
        "Authorization":        "token " + githubToken,
        "Accept":               "application/vnd.github+json",
        "X-GitHub-Api-Version": "2022-11-28",
    }
}

func GetUser(username string) (*models.User, error) {

    if user, err := getCachedUser(username); err == nil {
        return user, nil
    }

	url := fmt.Sprintf("https://api.github.com/users/%s", username)

	req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, fmt.Errorf("creating request: %w", err)
    }

	for k, v := range defaultHeaders {
        req.Header.Add(k, v)
    }

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

	var user models.User

	if err := json.Unmarshal(body, &user); err != nil {
        return nil, fmt.Errorf("parsing JSON: %w", err)
    }

    cacheUser(username, &user, 1800)

	return &user, nil
}

func GetUserRepos(repos_url string) ([]models.Repo, error) {

	url := repos_url

	req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, fmt.Errorf("creating request: %w", err)
    }

	for k, v := range defaultHeaders {
        req.Header.Add(k, v)
    }

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

	var repos []models.Repo

	if err := json.Unmarshal(body, &repos); err != nil {
        return nil, fmt.Errorf("parsing JSON: %w", err)
    }

	return repos, nil
}

func GetRepoContributors(repo_url string) ([]models.RepoContributors, error) {

	url := repo_url

	req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, fmt.Errorf("creating request: %w", err)
    }

	for k, v := range defaultHeaders {
        req.Header.Add(k, v)
    }

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

	var repoContributors []models.RepoContributors

	if err := json.Unmarshal(body, &repoContributors); err != nil {
        return nil, fmt.Errorf("parsing JSON: %w", err)
    }

	return repoContributors, nil
}