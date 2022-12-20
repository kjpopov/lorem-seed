// cobra.go
package cobra

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Url string
var Interval int
var PrintResponse bool

var rootCmd = &cobra.Command{
	Use:   "lorem-seed",
	Short: "Hits API endpoint to seed data or to perform load/stress test",
	Run: func(cmd *cobra.Command, args []string) {
		// code for parsing and handling command line arguments goes here
	},
}

func init() {
	rootCmd.Flags().StringVarP(&Url, "url", "u", "http://localhost:1323/api/todo", "URL to request")
	rootCmd.Flags().IntVarP(&Interval, "interval", "i", 10, "Interval between requests (in seconds)")
	rootCmd.Flags().BoolVarP(&PrintResponse, "print-response", "p", false, "Print response from server")
}

// ParseArguments parses the command line arguments
func ParseArguments() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
