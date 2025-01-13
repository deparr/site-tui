package tui

type asyncStatus int

const (
	empty asyncStatus = iota
	loading
	loaded
	errored
)

type asyncKey int

const (
	projectKey asyncKey = iota
)

type asyncJobMsg struct {
	key asyncKey
}

type asyncDoneMsg struct {
	key  asyncKey
	data map[string]any
	err  error
}

