package cmd

import (
	"github.com/frederikhs/planning-poker/app"
	"github.com/spf13/cobra"
	"log"
	"net/http"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve websocket server",
	Run: func(cmd *cobra.Command, args []string) {
		router := app.Create()
		log.Fatal(http.ListenAndServe(":3000", router))
	},
}
