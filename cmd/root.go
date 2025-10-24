package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var rootCmd = &cobra.Command{
	Use:   "dbtool",
	Short: "CLI tool for database management",
	Long:  "A tool for creating, deleting, migrating, and seeding the database.",
}

func loadEnvVariables() (string, string, string, string, string) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env")
	}

	user := os.Getenv("DB_USER")
	pwd := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	return user, pwd, host, port, name
}

func openDBConnection() (*gorm.DB, error) {
	user, pwd, host, port, name := loadEnvVariables()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", user, pwd, host, port, name)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func openSqlConnection() (*sql.DB, error) {
	user, pwd, host, port, _ := loadEnvVariables()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?parseTime=true&loc=Local", user, pwd, host, port)
	return sql.Open("mysql", dsn)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
