package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	version = "v0.5.15"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version information of Tunnel",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
