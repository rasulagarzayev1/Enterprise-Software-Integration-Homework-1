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
	"io/ioutil"
	"log"
	"encoding/json" 
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	
)

// todoitemCmd represents the todoitem command
var todoitemCmd = &cobra.Command{
	Use:   "todoitem",
	Short: "A one todo item",
	Long: `Our siple todo item which we want to see detailed`,
	Run: func(cmd *cobra.Command, args []string) {
		var id = args[0]
		getToDoItem(id)
	},
}

func init() {
	rootCmd.AddCommand(todoitemCmd)

}


func getToDoItem(id string){
	url := ("http://135.181.150.82:8080/api/v1/tasks/"+id)
	responseBytes := getToDoItemData(url)
	todo := Tasks{}

	if err := json.Unmarshal(responseBytes, &todo); err != nil {
		fmt.Printf("Could not unmarshal reponseBytes. %v", err)
	}

	if todo.Status == true {
		if todo.Level == 0 {
			color.Blue("[LOW]    " +string(todo.Title))
		}else if todo.Level == 1{
			color.Yellow("[MEDIUM] " + string(todo.Title))
		}else{
			color.Red("[HIGH]   " + string(todo.Title))
		}
	}else {
		if todo.Title == ""{
			color.Red("Not found")
		}else{
			color.Green("[COMPLETED]   " + string(todo.Title))
		}
	}

	

}

func getToDoItemData(baseAPI string) []byte{
	request, err := http.NewRequest(
		http.MethodGet,
		baseAPI,
		nil,
	)
	if err != nil {
		log.Printf("Could not request a list", err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "To-Do CLI")

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		log.Printf("Could not make a request", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body) 

	if err != nil {
		log.Printf("Could not make a request", err)
	}

	return responseBytes 
}