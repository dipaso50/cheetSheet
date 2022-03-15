/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"dipaso/cs/internal/application"
	"dipaso/cs/internal/domain"
	"dipaso/cs/internal/infraestructure"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cs",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {

		var allSheets []domain.Sheet

		if len(args) > 0 {
			param := args[0]
			allSheets = sheetService.List(param)
		} else {
			allSheets = sheetService.ListAll()
		}

		for _, sheet := range allSheets {
			fmt.Printf("%s\n", sheet.ID)
			fmt.Printf("Command: %s\n", sheet.Command)
			fmt.Printf("Description: %s\n\n", sheet.Description)

		}
	},
}

var sheetService application.SheetService

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	idGenerator := infraestructure.UUIDGenerator{}
	repo, err := infraestructure.NewJsonRepo(dirname, ".cheetSheet.json")

	if err != nil {
		log.Fatal(err)
	}

	sheetService = application.NewSheetService(idGenerator, repo)
}
