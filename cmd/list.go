package cmd

import (
	"fmt"
	"github.com/lucasnevespereira/lembra/internal/repository"
	"github.com/spf13/cobra"
)

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "List all reminders",
	Long:  "List displays all existing reminders",
	RunE:  listReminders,
}

func init() {
	rootCmd.AddCommand(listCommand)
}

func listReminders(cmd *cobra.Command, args []string) error {
	reminderRepo, err := repository.NewReminderRepository()
	if err != nil {
		return err
	}

	reminders, err := reminderRepo.GetAll()
	if err != nil {
		return err
	}

	// Print the reminders
	for _, r := range reminders {
		fmt.Printf("ID: %s\n", r.ID)
		fmt.Printf("Title: %s\n", r.Title)
		fmt.Printf("Message: %s\n", r.Message)
		fmt.Printf("Sound: %s\n", r.Sound)
		fmt.Printf("Time: %s\n", r.Time)
		fmt.Println("------")
	}

	return nil
}
