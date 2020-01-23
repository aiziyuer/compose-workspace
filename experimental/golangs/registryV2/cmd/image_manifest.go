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
			rootTree := tree.AddBranch(fmt.Sprintf("[ROOT] %s %d", ret.Digest, ret.Size))

			for _, m := range ret.Manifests {
				subTree := rootTree.AddBranch(fmt.Sprintf("[%s/%s] %s %d", m.Platform.Architecture, m.Platform.OS, m.Digest, m.Size))
				subTree.AddBranch(fmt.Sprintf("[ config] %s %d", m.Config.Digest, m.Config.Size))
				layerTree := subTree.AddBranch(fmt.Sprintf("[%7d] %s %d", len(m.Layers), "layers", m.Size-m.Config.Size))
				for i, layer := range m.Layers {
					layerTree.AddNode(fmt.Sprintf("[%3d] %s %d", i, layer.Digest, layer.Size))
				}
			}

			fmt.Println(tree.String())
			return nil
		},
	}

	imageCmd.AddCommand(imageManifestCmd)

}
