package cmd

import (
	"fmt"
	"github.com/lucasnevespereira/lembra/internal/pkg/repository"
	"github.com/lucasnevespereira/lembra/internal/pkg/repository/database"
	"github.com/lucasnevespereira/lembra/internal/utils/mapping"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
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

	db, err := database.Open()
	if err != nil {
		return fmt.Errorf("open db connection: %v\n", err)
	}
	reminderRepo := repository.NewReminderRepository(db)

	dbReminders, err := reminderRepo.GetAll()
	if err != nil {
		return err
	}
	reminders := mapping.ToRemindersDTO(dbReminders)

	table := tablewriter.NewWriter(os.Stdout)
	// Set the table header
	table.SetHeader([]string{"ID", "Title", "Message", "Time"})

	for _, r := range reminders {
		table.Append([]string{r.ID, r.Title, r.Message, r.Time})
	}

	table.Render()

	return nil
}
