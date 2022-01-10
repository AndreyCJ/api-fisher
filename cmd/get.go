package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get request",
	Long:  `Http GET request, supports Authorization token headers`,
	Run: func(cmd *cobra.Command, args []string) {
		GetRequest(cmd, args)
	},
}

func GetRequest(cmd *cobra.Command, args []string) {
	var URL string
	client := &http.Client{}

	token, _ := cmd.Flags().GetString(TOKEN_FLAG)

	fmt.Println("TOKEN => ", token)
	fmt.Println("Args => ", args)

	if len(args) == 0 {
		fmt.Println("Pass url to get the response")
		return
	} else {
		URL = args[0]
	}

	request, _ := http.NewRequest("GET", URL, nil)
	if token != "" {
		request.Header.Add("Authorization", "Bearer "+token)
		fmt.Println("REQ HEADER", request.Header)
	}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		formattedJSON, err := FormatJSON(response.Body)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("\n---")
		fmt.Println("Success: ")
		fmt.Println(string(formattedJSON))
		fmt.Println("\n---")

	} else {
		fmt.Println("\n---")
		fmt.Println("\nError: " + response.Status)
		fmt.Println("\n---")
	}
}

func init() {
	rootCmd.AddCommand(getCmd)
	UseTokenFlag(getCmd)
}