package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"lorem-seed/cobra"

	lr "github.com/UltiRequiem/lorelai/pkg"
)

func main() {
	cobra.ParseArguments()

	fmt.Printf("URL: %s\n", cobra.Url)
	fmt.Printf("Interval: %d\n", cobra.Interval)
	fmt.Printf("Print Response: %t\n", cobra.PrintResponse)

	// os.Exit(0)

	// Set the URL of the API endpoint
	url := cobra.Url
	interval := time.Duration(cobra.Interval) * time.Second

	log.Printf("Starting to seed %s with interval %s", url, interval)

	// Run the loop forever
	for {
		// Define the JSON payload to be sent with the PUT request
		payload := map[string]interface{}{
			"title":       strings.ToTitle(strings.TrimSpace(lr.LoremWords(5))),
			"description": lr.Sentence() + " " + lr.Sentence(),
			"completed":   false,
		}
		// Convert the payload to a JSON byte array
		payloadBytes, err := json.Marshal(payload)
		if err != nil {
			log.Println(err)
			break
		}

		// Create an HTTP request with the PUT method and the JSON payload
		req, err := http.NewRequest("PUT", url, bytes.NewBuffer(payloadBytes))
		if err != nil {
			fmt.Println(err)
			break
		}
		req.Header.Set("Content-Type", "application/json")

		// Send the request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
		} else {
			defer resp.Body.Close()
			if cobra.PrintResponse {
				// Read the response body
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					fmt.Println(err)
					return
				}
				printJSON(body)
			}
		}

		// Indicate the progress with a .
		if !cobra.PrintResponse {
			fmt.Print(".")
		}

		// Sleep for the specified interval before making the next request
		time.Sleep(interval)
	}
}
