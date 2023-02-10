package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "hello",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("yay")
	},
}
