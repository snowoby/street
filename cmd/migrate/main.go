package migrate

import (
	"log"
	"street/cmd/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {

	db := config.NewDefaultSql()
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("failed creating postgres driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"postgres", driver)
	if err != nil {
		log.Fatalf("failed creating migrate instance: %v", err)
	}

	err = m.Up()
	if err != nil {
		log.Fatalf("failed applying migrations: %v", err)
	}
	// migrate -source file://migrations -database mysql://root:pass@tcp(localhost:3306)/test up
}

func Main() {
	main()
}
