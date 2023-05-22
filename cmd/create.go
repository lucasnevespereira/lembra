package cmd

import (
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
	_ = createCommand.MarkFlagRequired("time")
}

var createCommand = &cobra.Command{
	Use:   "create",
	Short: "Create a new reminder",
	Long:  "Create creates a new reminder with a title message and given time",
	RunE:  createReminder,
}

func createReminder(cmd *cobra.Command, args []string) error {
	title, _ := cmd.Flags().GetString("title")
	message, _ := cmd.Flags().GetString("message")
	sound, _ := cmd.Flags().GetString("sound")
	timeStr, _ := cmd.Flags().GetString("time")

	reminderRepo, err := repository.NewReminderRepository()
	if err != nil {
		return err
	}

	time, err := reminder.ParseTime(timeStr)
	if err != nil {
		return err
	}
	newReminder, err := reminder.NewReminder(title, message, sound, time)
	if err != nil {
		return err
	}

	err = reminderRepo.Create(newReminder)
	if err != nil {
		return err
	}

	return nil
}
