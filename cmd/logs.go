package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"log"
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
		log.Fatalf("failed to get absolute path: %v", err)
	}

	file, err := os.Open(absPath)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading log file: %v", err)
	}
}
