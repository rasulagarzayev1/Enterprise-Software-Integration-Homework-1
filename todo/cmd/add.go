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
	"bytes"
    "encoding/json"
    "fmt"
    "net/http"
	"io"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adding todo",
	Long: `A longer description that user is creating new todo`,
	Run: func(cmd *cobra.Command, args []string){
		addNewItem()
	},
}
var Title string
var Status bool
var Level int
func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&Title, "Title", "t", "", "Title of the todo")
	addCmd.MarkFlagRequired("title")
	addCmd.Flags().BoolVarP(&Status, "Status", "s", true, "Status of todo")
	addCmd.Flags().IntVarP(&Level, "Level", "l", 0, "Level of todo")


}

func addNewItem(){
	var createdTask = Tasks {
		Title:        Title,
		Status:  	  Status,
		Level:		  Level,
	}

	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(createdTask)
	response, err := http.Post("http://135.181.150.82:8080/api/v1/tasks", "application/json", buffer)

	if err != nil {
		fmt.Println(err)
		color.Red("Something went wrong")
	}
	var createdT Tasks
	bytes, err := io.ReadAll(response.Body)

	json.Unmarshal(bytes, &createdT)

	if string(response.Status) == "201 Created" {
		color.Green("To do was created successfully")
	}else{
		color.Red("Couldnt create todo. Enter data correctly")
	}
	
}