package models

type User struct {
    Login       string `json:"login"`
    Name        string `json:"name"`
    AvatarURL   string `json:"avatar_url"`
    Bio         string `json:"bio"`
    Location    string `json:"location"`
    HTMLURL     string `json:"html_url"`
    CreatedAt   string `json:"created_at"`
    PublicRepos int    `json:"public_repos"`
	ReposUrl 	string `json:"repos_url"`
    Followers   int    `json:"followers"`
    Following   int    `json:"following"`
}


type Repo struct {
    Name        string `json:"name"`
    Description string `json:"description"`
	Language 	string `json:"language"`
    Stars       int    `json:"stargazers_count"`
    Forks       int    `json:"forks_count"`
    ContributorsURL string `json:"contributors_url"`
}

type RepoContributors struct {
    Name        string `json:"login"`
    HTMLURL string `json:"html_url"`
    Contributions   int `json:"contributions"`
}

type RepoWithContributors struct {
    RepoName     string
    Contributors []RepoContributors
}
