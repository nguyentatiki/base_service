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

var rootCmd = &cobra.Command{}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start the server",
	Run: func(cmd *cobra.Command, args []string) {
		protocol := os.Getenv("PROTOCOL")
		db := database.Open()
		container := internal.InitializeContainer(db)
		if protocol == "http" {
			http.Listen(":5000", container.HttpServer)
		} else {
			grpc.Register(":5001", container.GrpcServer)
		}
	},
}

func Execute() {
	rootCmd.AddCommand(startCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
