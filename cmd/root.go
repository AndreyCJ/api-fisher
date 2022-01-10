package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

const (
	TOKEN_FLAG = "token"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "api-fisher",
	Short: "A brief description of your application",
	Long:  `A small api tesrter`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
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