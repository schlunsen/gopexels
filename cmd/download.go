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
	"log"
	"sync"

	"github.com/cheggaaa/pb"
	"github.com/schlunsen/gopexels/pkg/helpers"
	"github.com/schlunsen/gopexels/pkg/pexels"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	NUMBER_OF_THEADS int
	NUMBER_OF_FILES  int
	OUTPUT_FOLDER    string
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download pexels objects",
	Long: `Download pexels objects:

`,
	Run: func(cmd *cobra.Command, args []string) {

		query := args[0]

		apiKey := viper.GetString("APIKEY")

		client := pexels.NewClient(apiKey)
		fmt.Println("download called", client.ApiKey)
		total := NUMBER_OF_FILES
		results := client.SimpleQuery(query, total)

		photos := results.Get("photos")

		threads := NUMBER_OF_THEADS

		downloadChan := make(chan string)
		stopChan := make(chan int)

		helpers.CreateDirIfNotExists(OUTPUT_FOLDER)

		var bar *pb.ProgressBar = pb.StartNew(total)
		var wg sync.WaitGroup

		for i := 0; i != threads; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for {
					select {
					case img := <-downloadChan:
						client.DownloadImage(img, OUTPUT_FOLDER)
						bar.Increment()
					case <-stopChan:
						return
					}
				}
			}()
		}

		for _, photo := range photos.MustArray() {
			ndata, _ := photo.(map[string]interface{})
			n2data, _ := ndata["src"].(map[string]interface{})
			downloadChan <- fmt.Sprintf("%s", n2data["original"])

		}

		// quit all the goroutines
		for i := 0; i != threads; i++ {
			stopChan <- 0
		}

		bar.Finish()
		log.Println("all images downloaded")
		log.Println("Waiting for the goroutines to exit...")
		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	downloadCmd.PersistentFlags().IntVar(&NUMBER_OF_THEADS, "threads", 3, "Number of threads to download with")
	downloadCmd.PersistentFlags().IntVar(&NUMBER_OF_FILES, "number", 10, "Number of files to download")
	downloadCmd.PersistentFlags().StringVar(&OUTPUT_FOLDER, "output", "output/", "Output folder name")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downloadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
