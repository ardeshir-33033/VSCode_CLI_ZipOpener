/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
	"uzo/utils"
)

// codeCmd represents the code command
var codeCmd = &cobra.Command{
	Use:   "code <zip_file_name.zip>",
	Short: "A brief description of your command",
	Long: `This Command Will help you to unzip and open your file in VS Code.
	In Order to work it is expected that VS Code is installed on your machine.`,
	Example: `uzo code SHIT.zip`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var fileName string
		var err error
		var argument string

		argument = args[0]
		fileExists, err := utils.FileExists(argument)
		if err != nil {
			fmt.Println(err.Error())
		}
		if fileExists {
			fileName, err = filepath.Abs(argument)
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			fmt.Printf("File %s does not exist", argument)
			return
		}

		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err.Error())
		}
		utils.Unzip(fileName, wd)

		os.Chdir(utils.FilenameWithoutExtension(fileName))

		commandCode := exec.Command("code", wd)
		err = commandCode.Run()

		if err != nil {
			fmt.Println("VS Code execution failed not found in %PATH%")
		}

	},
}

func init() {
	rootCmd.AddCommand(codeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// codeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// codeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
