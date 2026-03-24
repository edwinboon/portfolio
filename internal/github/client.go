package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/edwinboon/tui-portfolio/internal/models"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

type Repo struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	HTMLURL     string `json:"html_url"`
	Language    string `json:"language"`
}

const apiURL = "https://api.github.com/users/edwinboon/repos"

func FetchRepos() ([]Repo, error) {
	resp, err := httpClient.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("github api returned status %d", resp.StatusCode)
	}

	repos := []Repo{}
	err = json.NewDecoder(resp.Body).Decode(&repos)
	if err != nil {
		return nil, err
	}

	return repos, nil
}

func ToProject(r Repo) models.Project {
	return models.Project{
		Name:        r.Name,
		Description: r.Description,
		URL:         r.HTMLURL,
		Language:    r.Language,
	}
}
