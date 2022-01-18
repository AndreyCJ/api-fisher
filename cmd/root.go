package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

const (
	TOKEN_FLAG       = "token"
	INTERACTIVE_FLAG = "advanced"
	HEADERS_FLAG     = "headers"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "api-fisher",
	Short: "A brief description of your application",
	Long:  `A small api tesrter`,
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.api-fisher.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}

// GLOBAL FLAGS
func UseTokenFlag(cmd *cobra.Command) string {
	cmd.Flags().StringP(TOKEN_FLAG, "t", "", "auth token")
	return TOKEN_FLAG
}

func UseInteractiveFlag(cmd *cobra.Command) string {
	cmd.Flags().BoolP(INTERACTIVE_FLAG, "a", false, "advanced mode, allows input")
	return INTERACTIVE_FLAG
}

func UseHeadersFlag(cmd *cobra.Command) string {
	cmd.Flags().BoolP(HEADERS_FLAG, "h", false, "Allows headers input")
	return HEADERS_FLAG
}
