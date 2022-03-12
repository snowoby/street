package migrate

import (
	"context"
	"log"
	"street/cmd/config"

	"ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect/sql/schema"
)

func main() {
	ent := config.NewDefaultEnt()
	defer ent.Close()
	ctx := context.Background()
	dir, err := migrate.NewLocalDir("migrations")
	if err != nil {
		log.Fatalf("failed creating atlas migration directory: %v", err)
	}
	err = ent.Schema.Diff(ctx, schema.WithDir(dir))
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	// migrate -source file://migrations -database mysql://root:pass@tcp(localhost:3306)/test up
}

func Main() {
	main()
}
