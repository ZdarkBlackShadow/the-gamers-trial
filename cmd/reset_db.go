package cmd

func reset_db() {
	dropDb()
	create_db()
	migrate()
	seed()
}
