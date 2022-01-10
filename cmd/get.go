package cmd

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

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
	var payload *strings.Reader = nil

	token, _ := cmd.Flags().GetString(TOKEN_FLAG)
	isInteracitveMode, _ := cmd.Flags().GetBool(INTERACTIVE_FLAG)

	fmt.Println("TOKEN => ", token)
	fmt.Println("Args => ", args)

	if len(args) == 0 {
		fmt.Println("Pass url to get the response")
		return
	} else {
		URL = args[0]
	}

	if isInteracitveMode {
		fmt.Println("Enter request payload: ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			return
		}
		payload = strings.NewReader(input)

		fmt.Println("THE BODY IS => ", payload)
	}

	request, _ := http.NewRequest("GET", URL, payload)
	if token != "" {
		request.Header.Add("Authorization", "Bearer "+token)
		fmt.Println("REQ HEADER", request.Header)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		formattedJSON, err := FormatJSON(response.Body)

		if err != nil {
			fmt.Println(err)
		}

		PrintResults("Success: \n" + string(formattedJSON))
	} else {
		PrintResults("\nError: " + response.Status)
	}
}

func init() {
	rootCmd.AddCommand(getCmd)
	UseTokenFlag(getCmd)
	UseInteractiveFlag(getCmd)
}
