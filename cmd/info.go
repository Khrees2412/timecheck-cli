package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "get",
	Short: "To get the information you need",
	Long:  `This command is what helps timecheck know what you want to do`,
	Run: func(cmd *cobra.Command, args []string) {

		location, _ := cmd.Flags().GetString("location")
		if location != "" {
			data, err := FetchInfo(location)
			fmt.Println(PrettyPrint(data))
			if err != nil {
				fmt.Println(err)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(infoCmd)

	infoCmd.PersistentFlags().String("location", "", "The location you want to get info on.")
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
