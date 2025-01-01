package tui

import (
	"log/slog"
	"os"
)

// checks for vital env vars
func checkEnv() {
	for _, _var := range []string{
		"WEBSITE_URL",
		"GITHUB_URL",
		"GITHUB_USER",
		"LINKEDIN_USER",
		"API_URL",
		"EMAIL",
	} {
		_, ok := os.LookupEnv(_var)
		if !ok {
			slog.Warn("missing env", "var", _var)
		}
	}
}
