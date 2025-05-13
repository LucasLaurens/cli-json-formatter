package cmd

import (
	"cli-json-formatter/internal/service"
	"fmt"

	"github.com/spf13/cobra"
)

// formatCmd represents the format command
var formatCmd = &cobra.Command{
	Use:   "format",
	Short: "From String to Json",
	Long:  "Small command to transform String into Json format",
	Run: func(cmd *cobra.Command, args []string) {
		jsonAsString, _ := cmd.Flags().GetString("unparsed")
		formatted, err := service.FormatStringToJson(jsonAsString)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(formatted)
		}
	},
}

func init() {
	rootCmd.AddCommand(formatCmd)

	// Here you will define your flags and configuration settings.
	formatCmd.Flags().StringP("unparsed", "u", "", "unparsed json as string")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// formatCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// formatCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
