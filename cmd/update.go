package cmd

import (
	"fmt"
	"github.com/lucasnevespereira/lembra/internal/pkg/reminder"
	"github.com/lucasnevespereira/lembra/internal/pkg/repository"
	"github.com/lucasnevespereira/lembra/internal/pkg/repository/database"
	"github.com/spf13/cobra"
)

var updateCommand = &cobra.Command{
	Use:   "update",
	Short: "Update an existing reminder",
	Long:  "Update updates an existing reminder with the given ID",
	RunE:  updateReminder,
}

func init() {
	rootCmd.AddCommand(updateCommand)
	updateCommand.PersistentFlags().String("id", "", "Reminder ID")
	updateCommand.PersistentFlags().String("title", "", "Notification Title")
	updateCommand.PersistentFlags().String("message", "", "Notification Message")
	updateCommand.PersistentFlags().String("time", "", "Time of the reminder (format: 2006-01-02 15:04)")

	// Mark the required flags
	_ = updateCommand.MarkFlagRequired("id")
}

func updateReminder(cmd *cobra.Command, args []string) error {
	id, _ := cmd.Flags().GetString("id")
	title, _ := cmd.Flags().GetString("title")
	message, _ := cmd.Flags().GetString("message")
	timeStr, _ := cmd.Flags().GetString("time")

	db, err := database.Open()
	if err != nil {
		return fmt.Errorf("open db connection: %v\n", err)
	}
	reminderRepo := repository.NewReminderRepository(db)

	existingReminder, err := reminderRepo.GetByID(id)
	if err != nil {
		return fmt.Errorf("failed to get reminder: %v", err)
	}

	if title != "" {
		existingReminder.Title = title
	}
	if message != "" {
		existingReminder.Message = message
	}

	if timeStr != "" {
		time, err := reminder.ParseTime(timeStr)
		if err != nil {
			return err
		}
		existingReminder.Time = time
	}

	err = reminderRepo.Update(existingReminder)
	if err != nil {
		return err
	}

	return nil
}
