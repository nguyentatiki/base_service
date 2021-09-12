package cmd

import (
	"base_service/database"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
)

var migrateUpCmd *cobra.Command

func init() {
	migrateUpCmd = &cobra.Command{
		Use:   "up",
		Short: "migrate up command",
		Long:  "migrate up command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running migrate up command")
			db := database.Open()

			dbDriver, err := mysql.WithInstance(db.DB, &mysql.Config{})
			if err != nil {
				fmt.Printf("instance error: %v \n", err)
			}

			fileSource, err := (&file.File{}).Open("file://migrations")
			if err != nil {
				fmt.Printf("opening file error: %v \n", err)
			}

			m, err := migrate.NewWithInstance("file", fileSource, "mysql", dbDriver)
			if err != nil {
				fmt.Printf("migrate error: %v \n", err)
			}

			if err = m.Up(); err != nil {
				fmt.Printf("migrate up error: %v \n", err)
			}

			fmt.Println("Migrate up done with success")

		},
	}

	migrateCmd.AddCommand(migrateUpCmd)
}
