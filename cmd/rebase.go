/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

// rebaseCmd represents the rebase command
var rebaseCmd = &cobra.Command{
	Use:   "rebase [branch]",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		branch := args[0]
		out, _ := exec.Command("git", "branch", "--show-current").CombinedOutput()
		currentBranch := string(out[:len(out)-1])

		fmt.Printf("On branch %s. Rebasing on %s\n", currentBranch, branch)

		runCommand("git", "checkout", branch)

		fmt.Printf("Updating target branch %s\n", branch)
		runCommand("git", "pull", "origin", branch)

		fmt.Println("Target updated.")
		runCommand("git", "checkout", currentBranch)

		fmt.Printf("Rebasing %s\n", branch)
		runCommand("git", "rebase", branch)

		fmt.Printf("Rebase on branch %s ✅ successful", currentBranch)
	},
	Args: cobra.ExactArgs(1),
}

func runCommand(name string, args ...string) {
	c := exec.Command(name, args...)
	if out, err := c.CombinedOutput(); err != nil {
		log.Fatal(string(out))
	}
}

func init() {
	rootCmd.AddCommand(rebaseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rebaseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rebaseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
