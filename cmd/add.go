package cmd

import (
	"github.com/herrmannplatz/tilcli/internal/db"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func NewAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new learning",
		RunE: func(cmd *cobra.Command, args []string) error {
			l, err := db.Connect()
			if err != nil {
				return err
			}
			defer l.DB.Close()

			title, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Title").Show()
			desc, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Description").WithMultiLine().Show()

			if err := l.Insert(title, desc); err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}
