package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "gossaurus",
	Short: "Gossaurus is a rich dictionary generator for e-readers",
	Long:  "Gossaurus is a rich dictionary generator for e-readers see README.md for more details",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
