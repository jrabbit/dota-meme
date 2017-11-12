package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "meme",
	Short: " its a meme",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	cobra.OnInitialize(initConfig)
}

func Execute() {
	RootCmd.Execute()
}

func initConfig() {
	return
}
