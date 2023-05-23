package cmd

import (
	"fmt"
	"github.com/sevlyar/go-daemon"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func init() {
	rootCmd.AddCommand(stopCmd)
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops the reminder listener daemon",
	Run:   stopDaemon,
}

func stopDaemon(cmd *cobra.Command, args []string) {
	daemonContext := &daemon.Context{
		PidFileName: "lembra.pid",
		PidFilePerm: 0644,
		LogFileName: "lembra.log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
	}

	daemonPid, err := daemonContext.Search()
	if err != nil {
		log.Fatalf("failed to search for daemon process: %v", err)
	}

	err = daemonPid.Signal(os.Interrupt)
	if err != nil {
		log.Fatalf("failed to send termination signal to the daemon: %v", err)
	}

	err = daemonContext.Release()
	if err != nil {
		log.Fatalf("failed to release daemon resources: %v", err)
	}

	err = os.Remove(daemonContext.LogFileName)
	if err != nil {
		log.Fatalf("failed to delete lembra process logs: %v", err)
	}

	fmt.Println("Lembra Daemon stopped")
}
