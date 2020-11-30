package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:     "rm",
	Aliases: []string{"remove"},
	Short:   "Remove a task",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
		fmt.Println("rm called")
	},
}

func init() {
	RootCmd.AddCommand(rmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
