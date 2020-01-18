package cmd

import (
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

func init() {
	var imageCmd = &cobra.Command{
		Use:   "image",
		Short: "Image search/inspect ",
	}
	rootCmd.AddCommand(imageCmd)

	var imageSearchCmd = &cobra.Command{
		Use:   "search",
		Short: "SearchProject by name",
		RunE: func(cmd *cobra.Command, args []string) error {

			c, err := getClient()
			if err != nil {
				return err
			}

			// 业务逻辑
			projects, err := c.SearchProject("centos", -1)
			if err != nil {
				return err
			}

			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Name", "STARS", "OFFICIAL", "DESCRIPTION"})
			table.SetAlignment(tablewriter.ALIGN_LEFT)
			for _, v := range projects {
				table.Append(
					[]string{
						v.Name, strconv.Itoa(v.StartCount),
						map[bool]string{true: "Y", false: "N"}[v.Official],
						v.Description,
					},
				)
			}
			table.Render() // Send output

			return nil
		},
	}
	imageCmd.AddCommand(imageSearchCmd)

}
