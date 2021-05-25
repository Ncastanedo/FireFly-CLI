/*
Copyright © 2021 Kaleido, Inc.

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
	"errors"
	"fmt"

	"github.com/kaleido-io/firefly-cli/internal/stacks"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start <stack_name>",
	Short: "Start a stack",
	Long: `Start a stack

This command will start a stack and run it in the background.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("no stack specified")
		}
		stackName := args[0]

		stack, err := stacks.LoadStack(stackName)

		if err != nil {
			return err
		}

		if err = stack.StartStack(); err != nil {
			fmt.Printf("command finished with error: %v", err)
		} else {
			fmt.Print("\n\n")
			for _, member := range stack.Members {
				fmt.Printf("Web UI for member '%v': http://127.0.0.1:%v/ui\n", member.ID, member.ExposedFireflyPort)
			}
			fmt.Printf("\nTo see logs for your stack run:\n\nfirefly-cli logs %s\n\n", stackName)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
