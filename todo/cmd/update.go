/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"net/http"
	//"io/ioutil"
	"log"
	//"encoding/json" 
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Completing task",
	Long: `A longer description that  task is made completed`,
	Run: func(cmd *cobra.Command, args []string) {
		var id = args[0]
		updateItem(id)	
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)	
}

func updateItem(id string){
	url := fmt.Sprintf("http://135.181.150.82:8080/api/v1/tasks/"+id)

	request, err := http.NewRequest(
		http.MethodPut,
		url,
		nil,
	)
	if err != nil {
		log.Printf("Could not request a list", err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "To-Do CLI")

	response, err := http.DefaultClient.Do(request)
	if response.Status == "422 Unprocessable Entity" {
		color.Red("Could not make a request")
	}else{
		color.Green("Updated successfully")
	}

}

