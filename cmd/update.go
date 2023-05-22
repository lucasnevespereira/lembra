package cmd

import (
	"fmt"
	"github.com/lucasnevespereira/lembra/internal/reminder"
	"github.com/lucasnevespereira/lembra/internal/repository"
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
	updateCommand.PersistentFlags().String("sound", "", "Notification Sound")

	// Mark the required flags
	_ = updateCommand.MarkFlagRequired("id")
}

func updateReminder(cmd *cobra.Command, args []string) error {
	id, _ := cmd.Flags().GetString("id")
	title, _ := cmd.Flags().GetString("title")
	message, _ := cmd.Flags().GetString("message")
	sound, _ := cmd.Flags().GetString("sound")
	timeStr, _ := cmd.Flags().GetString("time")

	reminderRepo, err := repository.NewReminderRepository()
	if err != nil {
		return err
	}

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
	if sound != "" {
		existingReminder.Sound = sound
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
