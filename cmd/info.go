package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "get",
	Short: "To get the information you need",
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {

		location, _ := cmd.Flags().GetString("location")
		if location != "" {
			data, err := FetchInfo(location)
			fmt.Println(PrettyPrint(data))
			if err != nil {
				log.Fatalln(err)
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
