package cmd

import (
	"fmt"
	"github.com/frederikhs/planning-poker/app"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve websocket server",
	Run: func(cmd *cobra.Command, args []string) {
		router := app.Create()

		if os.Getenv("environment") == "production" {
			fmt.Println("running prod at 443")
			log.Fatal(http.ListenAndServeTLS(":443", "/etc/ssl/hrgn/tls.crt", "/etc/ssl/hrgn/tls.key", router))
		} else {
			log.Fatal(http.ListenAndServe(":3000", router))
		}
	},
}
