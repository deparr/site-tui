package tui

import (
	"log/slog"
	"os"
)

// checks for vital env vars
func checkEnv() {
	_, ok := os.LookupEnv("WEBSITE_URL")
	if !ok {
		slog.Warn("missing env", "var", "WEBSITE_URL")
	}

	_, ok = os.LookupEnv("GITHUB_URL")
	if !ok {
		slog.Warn("missing env", "var", "GITHUB_URL")
	}
}
