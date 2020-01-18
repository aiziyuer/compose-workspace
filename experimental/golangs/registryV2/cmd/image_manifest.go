package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

type ()

func init() {

	var imageManifestCmd = &cobra.Command{
		Use:   "manifest",
		Short: "manifest for image",
		RunE: func(cmd *cobra.Command, args []string) error {

			c, err := getClient()
			if err != nil {
				return err
			}

			// 业务逻辑
			manifests, err := c.ManifestV2(args[0])
			if err != nil {
				return err
			}

			fmt.Println(manifests)
			return nil
		},
	}

	imageCmd.AddCommand(imageManifestCmd)

}
