package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "format",
	Short: "From String to Json",
	Long:  "Small command to transform String into Json format",
	Run:   FormatStringToJson,
}

func FormatStringToJson(cmd *cobra.Command, args []string) {
	fmt.Println("format")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
