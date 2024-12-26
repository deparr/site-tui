package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/deparr/api/pkg/model"
)

type apiStatus int

const (
	statusOk apiStatus = 10
	missingUrl apiStatus = 11
)

var apiUrl string
var status apiStatus = statusOk

func Setup() {
	url, ok := os.LookupEnv("API_URL")
	if !ok {
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
	default:
		return "unknown or invalid status code"
	}
}

func unhealthyError() error {
	return fmt.Errorf("client: api unhealthy, %s", statusText(status))
}


// fetches repository data from the server.
// `which` is the kind of repos to fetch, can be "pinned" or "recent"
func GetRepos(which string) ([]model.Repository, error) {
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
		// should inspect error and determine if api server is somehow unhealthy
		return nil, err
	}
	defer res.Body.Close()

	var parsed struct {
		Data []model.Repository
	}
	err = json.NewDecoder(res.Body).Decode(&parsed)
	if err != nil {
		return nil, err
	}

	return parsed.Data, nil
}

