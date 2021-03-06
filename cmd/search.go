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

	"github.com/spf13/cobra"
)

var (
	city string
	url  string
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Searches a country for houses",
	Long: `Cobra is a CLI library for Go that empowers applications.
		   This application is a tool to generate the needed files
		   to quickly create a Cobra application.`,
	Run: searchCmdRun,
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringVarP(&city, "city", "c", "", "Crawl Houses in a city")
	searchCmd.Flags().StringVarP(&url, "url", "u", "", "The URL to crawl")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func searchCmdRun(cmd *cobra.Command, args []string) {
	for _, s := range args {
		fmt.Println("search Command:  " + s)
	}
	fmt.Println("City: " + city)
	fmt.Println("URL: " + url)
	crawl.crawl(url)
}
