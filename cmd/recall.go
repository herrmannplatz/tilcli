package cmd

import (
	"github.com/herrmannplatz/tilcli/internal/db"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func NewRecallCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "recall",
		Short: "Returns a random learning",
		RunE: func(cmd *cobra.Command, args []string) error {
			l, err := db.Connect()
			if err != nil {
				return err
			}
			defer l.DB.Close()

			learning, err := l.Random()
			if err != nil {
				return err
			}

			pterm.DefaultSection.Println(learning.Title)
			pterm.DefaultBox.WithTitle(learning.CreatedAt.Format("2006-01-02")).WithTitleBottomCenter().Println(learning.Description)

			return nil
		},
	}
	return cmd
}
