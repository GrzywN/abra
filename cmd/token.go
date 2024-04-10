package cmd

import (
	"fmt"
	"slices"

	"github.com/spf13/cobra"
)

var providers = []string{"deepl"}

var TokenCmd = &cobra.Command{
	Use:   "token",
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
	Run: token,
}

func init() {
  rootCmd.AddCommand(TokenCmd)
}

func token(cmd *cobra.Command, args []string) {
  fmt.Println(len(args))

  // provider := args[0]
  // key := args[1]
  //
  // if isValidProvider(provider) {
  //   fmt.Println("essa")
  //
  //   // ask for token
  // }
  //
  // fmt.Println(provider)
  // fmt.Println(key)
}

func isValidProvider(provider string) bool {
  return slices.Contains(providers, provider)
}
