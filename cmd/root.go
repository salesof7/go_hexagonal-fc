package cmd

import (
	"database/sql"
	"os"

	"github.com/salesof7/studies_go-hexagonal/adapters/db"
	"github.com/salesof7/studies_go-hexagonal/app"
	"github.com/spf13/cobra"
)

var Db, _ = sql.Open("sqlite3", "db.sqlite")
var productDb = db.NewProductDb(Db)
var productService = app.ProductService{Persistence: productDb}

var rootCmd = &cobra.Command{
	Use:   "studies_go-hexagonal",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
