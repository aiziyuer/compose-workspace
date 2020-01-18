package cmd

import (
	"errors"
	"fmt"
	"github.com/aiziyuer/registryV2/impl/util"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

func init() {

	var pageSize int
	var isShowTotalSize bool

	var imageSearchCmd = &cobra.Command{
		Use:   "search",
		Short: "SearchProject by name",
		RunE: func(cmd *cobra.Command, args []string) error {

			if len(args) < 1 {
				return errors.New("You need to provide words for search image! ")
			}

			if pageSize < -1 || pageSize == 0 {
				return errors.New("Page size should be natural number or -1 ! ")
			}

			c, err := getClient()
			if err != nil {
				return err
			}

			// 业务逻辑
			projects, err := c.SearchProject(args[0], pageSize)
			if err != nil {
				return err
			}

			// 结果输出
			switch outputFormat {
			case "table":
				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader([]string{"Name", "STARS", "OFFICIAL", "DESCRIPTION"})
				table.SetAlignment(tablewriter.ALIGN_LEFT)
				if isShowTotalSize {
					table.SetFooter([]string{"", "", "Total", strconv.Itoa(len(projects))})
				}
				table.SetBorder(true)
				for _, v := range projects {
					table.Append(
						[]string{
							v.Name,
							strconv.Itoa(v.StartCount),
							map[bool]string{true: "Y", false: "N"}[v.Official],
							v.Description,
						},
					)
				}
				table.Render()
			case "yaml":
				// TODO
			case "json":
				output := map[string]interface{}{
					"size": len(projects),
					"data": projects,
				}
				//json, err := util.Object2PrettyJson(output) // prettyjson 类库格式化有坑
				json, err := util.Object2Json(output)

				if err != nil {
					return err
				}
				fmt.Println(json)
			}

			return nil
		},
	}

	imageSearchCmd.PersistentFlags().IntVarP(
		&pageSize,
		"pageSize", "a", 10,
		"Page Size for search result, default 10, -1 for all",
	)

	imageSearchCmd.PersistentFlags().BoolVarP(
		&isShowTotalSize,
		"show-total", "t", false,
		"Show total size at the footer",
	)

	imageCmd.AddCommand(imageSearchCmd)

}
