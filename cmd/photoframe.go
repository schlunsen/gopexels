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
	"github.com/schlunsen/gopexels/pkg/server"
	"github.com/spf13/cobra"
	"github.com/webview/webview"
)

// photoframeCmd represents the photoframe command
var photoframeCmd = &cobra.Command{
	Use:   "photoframe",
	Short: "Starts a slideshow",
	Long:  `Starts a slideshow`,
	Run: func(cmd *cobra.Command, args []string) {
		query := ""
		if len(args) > 0 {
			query = args[0]
		}

		debug := false
		w := webview.New(debug)

		defer w.Destroy()
		w.SetTitle("Minimal webview example")
		w.SetSize(1980, 1080, webview.HintNone)

		w.Navigate("http://localhost:8080/start")
		go server.StartServer(query)
		w.Run()

	},
}

func init() {
	rootCmd.AddCommand(photoframeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// photoframeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// photoframeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
