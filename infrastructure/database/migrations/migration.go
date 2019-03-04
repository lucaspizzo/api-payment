package migrations

import (
	"log"

	db "github.com/lucaspizzo/api-payment/infrastructure/database"
	"gopkg.in/gormigrate.v1"
)

var migrations = []*gormigrate.Migration{&migration201903041722}

func RunMigrations() {
	postgres := &db.Repository{}
	db := postgres.GetInstance()
	m := gormigrate.New(db, gormigrate.DefaultOptions, migrations)
	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration did run successfully")
}
