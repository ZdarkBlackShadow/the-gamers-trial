package cmd

import "github.com/spf13/cobra"

var createDbCmd = &cobra.Command{
	Use:   "create-db",
	Short: "Create the database",
	Run: func(cmd *cobra.Command, args []string) {
		create_db()
	},
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Apply GORM migration",
	Run: func(cmd *cobra.Command, args []string) {
		migrate()
	},
}

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Insert default data",
	Run: func(cmd *cobra.Command, args []string) {
		seed()
	},
}

var dropDbCmd = &cobra.Command{
	Use:   "drop-db",
	Short: "Delete database",
	Run: func(cmd *cobra.Command, args []string) {
		dropDb()
	},
}

var resetDbCmd = &cobra.Command{
	Use:   "reset-db",
	Short: "Delete the database and Recreate the database",
	Run: func(cmd *cobra.Command, args []string) {
		reset_db()
	},
}

func init() {
	rootCmd.AddCommand(createDbCmd)
	rootCmd.AddCommand(migrateCmd)
	rootCmd.AddCommand(seedCmd)
	rootCmd.AddCommand(dropDbCmd)
	rootCmd.AddCommand(resetDbCmd)
}
