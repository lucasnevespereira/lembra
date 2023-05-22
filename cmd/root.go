package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "lembra",
	Short: "lembra is a reminder cli tool",
	Long:  "lembra is a reminder cli tool that allows to set up a reminder in a given time, that reminder executes a OS notification",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
