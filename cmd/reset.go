package cmd

import (
	"github.com/herrmannplatz/tilcli/internal/db"
	"github.com/spf13/cobra"
)

func NewResetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reset",
		Short: "Delete all learnings",
		RunE: func(cmd *cobra.Command, args []string) error {
			l, err := db.Connect()
			if err != nil {
				return err
			}
			defer l.DB.Close()

			if err := l.Reset(); err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}
