package client

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/deparr/api/model"
)

// todo consider making the client a struct instead of package level

type apiStatus int

const (
	statusOk          apiStatus = 10
	missingUrl        apiStatus = 11
	serverUnreachable apiStatus = 12

	unhealthyThreshold int = 3
)

var apiUrl string
var status apiStatus = statusOk

func Init() {
	url, ok := os.LookupEnv("API_URL")
	if !ok {
		slog.Warn("client: no api url given, dynamic content unavailable")
		status = missingUrl
	} else {
		apiUrl = url
	}
}

func ApiIsHealthy() bool {
	return isHealthyStatus()
}

func isHealthyStatus() bool {
	return status <= statusOk
}

func statusText(s apiStatus) string {
	switch s {
	case statusOk:
		return "OK"
	case missingUrl:
		return "missing or invalid url"
	case serverUnreachable:
		return "passed health check threshold, server unreachable"
	default:
		return "unknown or invalid status code"
	}
}

func unhealthyError() error {
	return fmt.Errorf("client: api unhealthy, %s", statusText(status))
}

func GetGhPinned() ([]model.Repository, error) {
	return getRepos("pinned")
}

func GetGhRecent() ([]model.Repository, error) {
	return getRepos("recent")
}

// fetches repository data from the server.
// `which` is the kind of repos to fetch, can be "pinned" or "recent"
func getRepos(which string) ([]model.Repository, error) {

	if !isHealthyStatus() {
		return nil, unhealthyError()
	}

	// todo better way to enforce correct routes
	if which != "pinned" && which != "recent" {
		return nil, fmt.Errorf("client.GetRepos: unknown repo type '%s'", which)
	}

	reqUrl := fmt.Sprintf("%s%s%s", apiUrl, "/gh/", which)
	res, err := http.Get(reqUrl)
	if err != nil {
		// todo should inspect error and determine if api server is somehow unhealthy
		// maybe factor into a `do` method once more routes are used
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("client.GetRepos: api bad status %d", res.StatusCode)
	}

	defer res.Body.Close()

	var parsed struct{ Data []model.Repository }
	err = json.NewDecoder(res.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}

	return parsed.Data, nil
}
