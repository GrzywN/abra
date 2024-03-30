package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "abra",
	Short: "An application for adding translated entries in ARB files.",
	Long: `
Abra is a CLI tool mainly for Flutter developers
that makes adding ARB entries to your application much easier.
Provide tokens for translation services and add translations using Abra CLI.

Usage:

abra init - Initializes config files for abra in your project root.
abra token deepl <token> - Sets up an API key for selected provider (DeepL)
abra translate - Add translated entry for your arb files
`,
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.abra.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

