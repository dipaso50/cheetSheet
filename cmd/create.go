/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var commandName string
var commandDescription string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		idresult := sheetService.CreateSheet(commandName, commandDescription)
		fmt.Printf("%s created\n", idresult)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&commandName, "name", "n", "", "the command that we want to create")
	createCmd.Flags().StringVarP(&commandDescription, "description", "d", "", "the command description")

	createCmd.MarkFlagRequired("name")
	createCmd.MarkFlagRequired("description")
}
