package cmd

import (
	"fmt"
	"os"

	"base_service/database"
	"base_service/internal"
	"base_service/internal/api/grpc"
	"base_service/internal/api/http"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "start",
	Short: "start the server",
	Run: func(cmd *cobra.Command, args []string) {
		protocol := os.Getenv("PROTOCOL")
		if protocol == "http" {
			http.Listen(":5000")
		} else {
			db := database.Open()
			//conn := os.Getenv("MYSQL_CONNECTION")
			//db, err := sqlx.Connect("mysql", conn)
			// if err != nil {
			// 	panic(err)
			// }
			s := internal.InitializeContainer(db)
			grpc.Register(":5001", s)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
