/*
Copyright © 2024 Vinicius Gregorio - vincamgreg@hotmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vinicius-gregorio/go_stress/cmd/external"
	"github.com/vinicius-gregorio/go_stress/internal"
)

// stressTestCmd represents the stressTest command
var stressTestCmd = &cobra.Command{
	Use:   "stressTest",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if url != "" {

			external.PrintBox(fmt.Sprintf("Stress Testing URL: %s", url))

		} else {
			fmt.Println("StressTest command: No URL provided for stress test")
		}
		err := validateFlags()
		if err != nil {
			fmt.Println(err)
		}
		st, err := internal.NewStressTest(url, requestCount, concurrency)
		if err != nil {
			fmt.Println(err)
		}
		st.Run()

	},
}

func init() {
	rootCmd.AddCommand(stressTestCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stressTestCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stressTestCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func validateFlags() error {
	if url == "" {
		return fmt.Errorf("URL is required")
	}
	if requestCount == 0 {
		return fmt.Errorf("request Count is required")
	}
	if concurrency == 0 {
		return fmt.Errorf("concurrency is required")
	}
	return nil
}
