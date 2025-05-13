/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
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

		formatted, err := FormatStringToJson(jsonAsString)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(formatted)
		}
	},
}

// todo: Cut into several sub methods and files
func FormatStringToJson(value string) (string, error) {
	// todo: create a struct to replace basic one
	var formattedJson json.RawMessage
	err := json.Unmarshal([]byte(value), &formattedJson)
	if err != nil {
		return "", fmt.Errorf(
			"you cannot parse this string to json format: %v",
			err,
		)
	}

	indentedJson, err := json.MarshalIndent(formattedJson, "", "  ")
	if err != nil {
		return "", fmt.Errorf(
			"you are not allowed to indent the json: %v",
			err,
		)
	}

	// todo: export formatted json as a file
	return string(indentedJson), nil
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
