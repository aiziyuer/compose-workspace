package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	var imageCmd = &cobra.Command{
		Use:   "image",
		Short: "Image search/inspect ",
	}
	rootCmd.AddCommand(imageCmd)

	var imageSearchCmd = &cobra.Command{
		Use:   "search",
		Short: "Search by name",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("image called")

			return errors.New("")
		},
	}
	imageCmd.AddCommand(imageSearchCmd)

}
