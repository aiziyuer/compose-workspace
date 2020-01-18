package cmd

import (
	"github.com/spf13/cobra"
)

var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "Image search/inspect ",
}

func init() {
	rootCmd.AddCommand(imageCmd)
}
