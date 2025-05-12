package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "format",
	Short: "From String to Json",
	Long:  "Small command to transform String into Json format",
	RunE:  FormatStringToJson,
}

// todo: Cut into several sub methods and files
func FormatStringToJson(cmd *cobra.Command, args []string) error {
	fmt.Println("format string to json", reflect.TypeOf(args[0]))

	if len(args) != 1 {
		return errors.New("you have to get the json as string argument")
	}

	jsonAsString := args[0]
	var formattedJson json.RawMessage
	err := json.Unmarshal([]byte(jsonAsString), &formattedJson)
	if err != nil {
		return fmt.Errorf(
			"you cannot parse this string to json format: %v",
			err,
		)
	}

	indentedJson, err := json.MarshalIndent(formattedJson, "", "  ")
	if err != nil {
		return fmt.Errorf(
			"you are not allowed to indent the json: %v",
			err,
		)
	}

	fmt.Println(string(indentedJson))
	// todo: export formatted json as a file

	return nil
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
