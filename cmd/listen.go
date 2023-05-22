package cmd

import (
	"github.com/lucasnevespereira/lembra/internal/notifier"
	"github.com/sevlyar/go-daemon"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	rootCmd.AddCommand(listenCmd)
}

var listenCmd = &cobra.Command{
	Use:   "listen",
	Short: "Starts the reminder listener daemon",
	Run:   runDaemon,
}

func runDaemon(cmd *cobra.Command, args []string) {
	daemonContext := &daemon.Context{
		PidFileName: "lembra.pid",
		PidFilePerm: 0644,
		LogFileName: "lembra.log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
	}

	child, err := daemonContext.Reborn()
	if err != nil {
		log.Fatalf("failed to start the daemon: %v", err)
	}
	if child != nil {
		os.Exit(0)
	}
	defer daemonContext.Release()

	// Start the notifier
	notifier, err := notifier.NewCronNotifier()
	if err != nil {
		log.Fatalf("creating notifier: %v", err)
	}

	if err := notifier.Start(); err != nil {
		log.Fatalf("starting notifier: %v", err)
	}

	// Wait for termination signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan

	// Stop the notifier
	notifier.Stop()
}
