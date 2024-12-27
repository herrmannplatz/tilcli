package cmd

import (
	"fmt"

	"github.com/herrmannplatz/tilcli/internal/db"
	"github.com/spf13/cobra"
)

func NewListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all learnings",
		RunE: func(cmd *cobra.Command, args []string) error {
			l, err := db.Connect()
			if err != nil {
				return err
			}
			defer l.DB.Close()

			learnings, err := l.GetAll()
			if err != nil {
				return err
			}
			for _, learning := range learnings {
				fmt.Println(learning.Description)
				fmt.Println()
			}
			return nil
		},
	}
	return cmd
}
