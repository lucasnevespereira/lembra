package cmd

import (
	"fmt"
	"github.com/lucasnevespereira/lembra/internal/reminder"
	"github.com/lucasnevespereira/lembra/internal/repository"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCommand)
	createCommand.PersistentFlags().String("title", "Lembra", "Notification Title")
	createCommand.PersistentFlags().String("message", "", "Notification Message")
	createCommand.PersistentFlags().String("sound", "default", "Notification Sound")
	createCommand.PersistentFlags().String("time", "", "Time of the reminder (format: 2006-01-02 15:04)")

	// Mark the required flags
	//_ = createCommand.MarkFlagRequired("message")
	//_ = createCommand.MarkFlagRequired("time")
}

var createCommand = &cobra.Command{
	Use:   "create",
	Short: "Create a new reminder",
	Long:  "Create creates a new reminder with a title message and given time",
	RunE: func(cmd *cobra.Command, args []string) error {
		title, _ := cmd.Flags().GetString("title")
		message, _ := cmd.Flags().GetString("message")
		sound, _ := cmd.Flags().GetString("sound")
		timeStr, _ := cmd.Flags().GetString("time")

		// Create a new ReminderRepository with the database connection
		reminderRepo, err := repository.NewReminderRepository()
		if err != nil {
			return err
		}
		defer reminderRepo.Close()

		// Create a new Reminder instance
		time, err := reminder.ParseTime(timeStr)
		if err != nil {
			return fmt.Errorf("failed to parse reminder time: %v", err)
		}
		newReminder := reminder.NewReminder(title, message, sound, time)

		// Save the reminder using the repository
		err = reminderRepo.Save(newReminder)
		if err != nil {
			return fmt.Errorf("failed to save reminder: %v", err)
		}

		return nil
	},
}
