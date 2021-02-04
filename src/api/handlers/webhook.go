package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"
)

type Owner struct {
	Login     string `json:"login"`
	Type      string `json:"type"`
	SiteAdmin bool   `json:"site_admin"`
}

type Repository struct {
	Url   string `json:"url"`
	Id    int    `json:"id"`
	Owner `json:"owner"`
}

type PullRequest struct {
	Number int `json:"number"`
	Title string `json:"title"`

}

type Sender struct {
	Login     string `json:"login"`
}

type GitHub struct {
	Action      string `json:"action"`
	Number      uint   `json:"number"`
	Repository  `json:"repository"`
	PullRequest `json:"pull_request,omitempty"`
	Sender      `json:"sender"`
}

func WebHook(c echo.Context) error {
	defer c.Request().Body.Close()
	var gitHub GitHub
	err := json.NewDecoder(c.Request().Body).Decode(&gitHub)

	gitHub.PrintGitHubPR()
	if err != nil {
		log.Printf("Failed processing Person request: %s\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	response := fmt.Sprintf("PullRequest Title \"%s\"", gitHub.PullRequest.Title)

	c.Response().Write([]byte(response))
	return nil
}

func (gitHub *GitHub) PrintGitHubPR() {
	log.Printf("%+v\n", gitHub)
}
