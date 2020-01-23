package cmd

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/xlab/treeprint"
)

func init() {

	var imageManifestCmd = &cobra.Command{
		Use:   "manifest",
		Short: "manifest for image",
		RunE: func(cmd *cobra.Command, args []string) error {

			if len(args) < 1 {
				return errors.New("input image name first ")
			}

			c, err := getClient()
			if err != nil {
				return err
			}

			// 业务逻辑
			ret, err := c.ManifestV2(args[0])
			if err != nil {
				logrus.Errorf("can't get ret from image(%s)", args[0])
				return err
			}

			tree := treeprint.New()
			rootTree := tree.AddBranch(fmt.Sprintf("[D] %s %d", ret.Digest, ret.Size))

			for _, m := range ret.Manifests {
				subTree := rootTree.AddBranch(fmt.Sprintf("[P %s/%s] %s %d", m.Platform.OS, m.Platform.Architecture, m.Digest, m.Size))
				subTree.AddNode(fmt.Sprintf("[C] %s %d", m.Config.Digest, m.Config.Size))
				//size := len(m.Layers)
				for i, layer := range m.Layers {
					subTree.AddNode(fmt.Sprintf("[L %3d] %s %d", i+1, layer.Digest, layer.Size))
				}
			}

			fmt.Println(tree.String())
			return nil
		},
	}

	imageCmd.AddCommand(imageManifestCmd)

}
