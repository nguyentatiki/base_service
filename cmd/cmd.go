package cmd

import (
	"fmt"
	"os"

	"base_service/internal/api/grpc"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "start",
	Short: "start the server",
	Run: func(cmd *cobra.Command, args []string) {
		grpc.Register(":5001")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
