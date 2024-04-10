package token

import (
	"fmt"
	"slices"
  "abra/cmd"

	"github.com/spf13/cobra"
)

var providers = []string{"deepl"}

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: set,
}

func init() {
	cmd.TokenCmd.AddCommand(setCmd)
}

func set(cmd *cobra.Command, args []string) {
  // TODO: dodać zabezpieczenie dla dlugosci args

  if len(args) < 2 {
    fmt.Println("Usage: abra token set <provider> <token>")

    return
  }

  provider := args[0]
  key := args[1]

  if isValidProvider(provider) {
    fmt.Println("provider is valid")

    return;
  }

  fmt.Printf(key)

  return;
}

func isValidProvider(provider string) bool {
  return slices.Contains(providers, provider)
}
