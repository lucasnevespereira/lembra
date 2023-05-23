package cmd

import (
	"bufio"
	"fmt"
	"github.com/lucasnevespereira/lembra/internal/utils/logger"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

func init() {
	rootCmd.AddCommand(readLogCmd)
}

var readLogCmd = &cobra.Command{
	Use:   "logs",
	Short: "Reads the lembra log file",
	Run:   readLog,
}

func readLog(cmd *cobra.Command, args []string) {
	logFilePath := "lembra.log"
	absPath, err := filepath.Abs(logFilePath)
	if err != nil {
		logger.Log.Fatalf("failed to get absolute path: %v", err)
	}

	file, err := os.Open(absPath)
	if err != nil {
		logger.Log.Fatalf("failed to open log file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		logger.Log.Fatalf("error reading log file: %v", err)
	}
}
