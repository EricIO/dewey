package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "dewey",
		Short: "A personal library manager",
		Long: `dewey is a software program that enables you to
manage your personal library.`,
		Version: "0.1.0",
	}
)

func Execute() error {
	return rootCmd.Execute()
}
