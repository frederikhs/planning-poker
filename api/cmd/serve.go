package cmd

import (
	"github.com/frederikhs/planning-poker/app"
	"github.com/spf13/cobra"
	"log"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve websocket server",
	Run: func(cmd *cobra.Command, args []string) {
		log.Fatal(app.Create().Listen(":3000"))
	},
}
