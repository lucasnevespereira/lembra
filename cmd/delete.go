package cmd

import (
	"fmt"
	"github.com/lucasnevespereira/lembra/internal/pkg/storage"
	"github.com/lucasnevespereira/lembra/internal/utils/logger"
	"github.com/spf13/cobra"
)

var deleteCommand = &cobra.Command{
	Use:   "delete",
	Short: "Delete a reminder",
	Long:  "Delete deletes an existing reminder with the given ID",
	RunE:  deleteReminder,
}

func init() {
	rootCmd.AddCommand(deleteCommand)
	deleteCommand.PersistentFlags().String("id", "", "Reminder ID")
	deleteCommand.PersistentFlags().Bool("all", false, "Delete all reminders")
}

func deleteReminder(cmd *cobra.Command, args []string) error {
	id, _ := cmd.Flags().GetString("id")
	deleteAll, _ := cmd.Flags().GetBool("all")

	dbFile, err := storage.OpenStorageFile()
	if err != nil {
		return err
	}
	reminderRepo := storage.NewReminderStorage(dbFile)

	if deleteAll {
		err := reminderRepo.DeleteAll()
		if err != nil {
			return err
		}

		logger.Log.Infoln("All reminders deleted successfully")
	} else {
		existingReminder, err := reminderRepo.GetByID(id)
		if err != nil {
			return err
		}

		if existingReminder.ID != "" {
			err = reminderRepo.DeleteByID(existingReminder.ID)
			if err != nil {
				return err
			}

			fmt.Println("Reminder deleted successfully")
		}
	}

	return nil
}
