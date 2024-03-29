package cmd

import (
	"github.com/lucasnevespereira/lembra/internal/pkg/notifier"
	"github.com/lucasnevespereira/lembra/internal/pkg/storage"
	"github.com/lucasnevespereira/lembra/internal/utils/logger"
	"github.com/robfig/cron"
	"github.com/sevlyar/go-daemon"
	"github.com/spf13/cobra"
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
		logger.Log.Fatalf("failed to start the daemon: %v", err)
	}
	if child != nil {
		os.Exit(0)
	}
	defer func(daemonContext *daemon.Context) {
		err := daemonContext.Release()
		if err != nil {
			logger.Log.Errorf("failed to release daemon resources: %v", err)
		}
	}(daemonContext)

	// Perform DB operations
	dbFile, err := storage.OpenStorageFile()
	if err != nil {
		logger.Log.Fatalf("open storage file: %v", err)
	}
	reminderRepo := storage.NewReminderStorage(dbFile)
	notifier := notifier.NewCronNotifier(reminderRepo, cron.New())
	if err := notifier.Start(); err != nil {
		logger.Log.Fatalf("starting notifier: %v", err)
	}

	// Wait for termination signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan

	notifier.Stop()
}
