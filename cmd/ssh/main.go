package main

import (
	"context"
	"errors"
	"net"
	"os"
	"os/signal"
	"syscall"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"
	"github.com/deparr/portfolio/go/pkg/tui"
	"github.com/joho/godotenv"
	"github.com/muesli/termenv"
	gossh "golang.org/x/crypto/ssh"
)

const (
	host = "0.0.0.0"
	port = "22"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Error("godotenv:", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signalChan
		cancel()
	}()
	s, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(host, port)),
		wish.WithHostKeyPath(".ssh/id_ed25519"),
		wish.WithMiddleware(
			bubbletea.Middleware(teaHandler),
			activeterm.Middleware(), // Bubble Tea apps usually require a PTY.
			logging.Middleware(),
		),
		wish.WithPublicKeyAuth(func(_ ssh.Context, key ssh.PublicKey) bool {
			return true
		}),
		wish.WithKeyboardInteractiveAuth(func(ctx ssh.Context, challenger gossh.KeyboardInteractiveChallenge) bool {
			return true
		}),
	)
	if err != nil {
		log.Error("Could not create server", "error", err)
	}

	log.Info("Starting SSH server", "port", port)
	go func() {
		if err = s.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
			log.Error("Could not start server", "error", err)
			cancel()
		}
	}()

	<-ctx.Done()
	s.Shutdown(ctx)
	log.Info("Stopping SSH server")
	cancel()
}

// TODO: this bridge shouldn't be needed anymore, but I still can't get the
//  colors to work correctly when run as a systemd service
//	01/06/24 -- this bridge is most definitely needed, otherwise firest
//	render is delayed and colors are off. It shouldn't be needed from
//	everything I've read but it is, so

type sshOutput struct {
	ssh.Session
	tty *os.File
}

func (s *sshOutput) Write(p []byte) (int, error) {
	return s.Session.Write(p)
}

func (s *sshOutput) Read(p []byte) (int, error) {
	return s.Session.Read(p)
}

func (s *sshOutput) Fd() uintptr {
	return s.tty.Fd()
}

func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	pty,_,_ := s.Pty()
	sshPty := &sshOutput{
		Session: s,
		tty:     pty.Slave,
	}

	renderer := bubbletea.MakeRenderer(sshPty)
	// this just doesn't do anything apparently
	renderer.SetColorProfile(termenv.TrueColor)
	model := tui.NewModel(renderer)
	return model, []tea.ProgramOption{tea.WithAltScreen()}
}
