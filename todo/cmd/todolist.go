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
	//"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json" 
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	
)

// todolistCmd represents the todolist command
var todolistCmd = &cobra.Command{
	Use:   "todolist",
	Short: "This is a list which show todos",
	Long: `A longer description that list not only show tasks also show the privority levels of our todos with colours.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("todolist called")
		getToDoList()
	},
}

func init() {
	rootCmd.AddCommand(todolistCmd)
}


func getToDoList(){
	url := "http://135.181.150.82:8080/api/v1/tasks"
	responseBytes := getToDoListData(url)
	var todo []Tasks
	var completed []Tasks
	var inProgress []Tasks

	if err := json.Unmarshal(responseBytes, &todo); err != nil {
		color.Red("Could not unmarshal reponseBytes. %v", err)
	}
	
	for i := 0; i < len(todo); i++ {
		if todo[i].Status == false{
			completed = append(completed, todo[i])
		}else{
			inProgress = append(inProgress, todo[i])
		}
		
	}
	for i := 0; i < len(inProgress); i++ {
		if inProgress[i].Level == 0 {
			color.Blue("[LOW]    " +string(inProgress[i].Title))
		}else if inProgress[i].Level == 1{
			color.Yellow("[MEDIUM] " + string(inProgress[i].Title))
		}else{
			color.Red("[HIGH]   " + string(inProgress[i].Title))
		}
	}
	for i := 0; i < len(completed); i++ {
		color.Green("[COMPLETED] " +string(completed[i].Title))
	}


}

func getToDoListData(baseAPI string) []byte{
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

