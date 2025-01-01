package client

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/deparr/api/pkg/model"
)

type apiStatus int

const (
	statusOk          apiStatus = 10
	missingUrl        apiStatus = 11
	serverUnreachable apiStatus = 12

	unhealthyThreshold int = 3
)

var apiUrl string
var status apiStatus = statusOk

var badHealthChecks int
var healthCheckTicker *time.Ticker

func Init() {
	url, ok := os.LookupEnv("API_URL")
	if !ok {
		slog.Warn("client: no api url given, dynamic content unavailable")
		status = missingUrl
	} else {
		apiUrl = url
	}

	healthCheckTicker = time.NewTicker(20 * time.Second)
	badHealthChecks = 0
	// TODO dont do this, just try to load at startup and then allow a manual rety on errors
	// go func() {
	// 	hcUrl := apiUrl + "/health"
	// 	for {
	// 		select {
	// 		case _ = <-healthCheckTicker.C:
	// 			slog.Info("ping")
	// 			res, err := http.Get(hcUrl)
	// 			defer res.Body.Close()
	// 			if err != nil {
	// 				slog.Error("client.HealthCheck", "err", err, "status", res.Status)
	// 			}
	//
	// 			if res.StatusCode != http.StatusOK {
	// 				badHealthChecks += 1
	// 				slog.Warn("failed server health check", "count", badHealthChecks)
	// 				if badHealthChecks >= unhealthyThreshold {
	// 					slog.Error("passed health check threshold")
	// 					status = serverUnreachable
	// 				}
	// 			} else {
	// 				badHealthChecks = 0
	// 			}
	// 		}
	// 	}
	// }()
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
