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

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:              "status",
	TraverseChildren: true,
	Short:            "status your branch onto another branch",
	Long: `A longer description that spans multiple lines and likely contains examples
			and usage of using your command. For example:
			
			Cobra is a CLI library for Go that empowers applications.
			This application is a tool to generate the needed files
			to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {
		long, err := cmd.Flags().GetBool("long")
		if err != nil {
			log.Fatal(err)
		}

		var out []byte

		if long {
			out, _ = exec.Command("git", "status").Output()
		}

		out, _ = exec.Command("git", "status", "-s").Output()

		fmt.Println("status called\n", string(out))
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	statusCmd.PersistentFlags().String("foo", "", "A help for foo")
	statusCmd.PersistentFlags().BoolP("long", "l", false, "longhand format")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
