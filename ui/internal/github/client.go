package github

import (
	"encoding/json"
	"net/http"
)

type Repo struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	HTMLURL     string `json:"html_url"`
	Language    string `json:"language"`
}

const apiURL = "https://api.github.com/users/edwinboon/repos"

func FetchRepos() ([]Repo, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}

	// Ensure the response body is closed after we're done with it
	defer resp.Body.Close()

	repos := []Repo{}
	err = json.NewDecoder(resp.Body).Decode(&repos)
	if err != nil {
		return nil, err
	}

	return repos, nil
}
