package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/frianlh/pokedex-api/migrations"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

const (
	Up    = "up"
	Down  = "down"
	Force = "force"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds)

	// migration config
	newMigrationConfig := migrations.NewMigrationConfig()
	migrationConfig, err := newMigrationConfig.Read()
	if err != nil {
		log.Fatal(err)
		return
	}
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		migrationConfig.PostgresConfig.User,
		migrationConfig.PostgresConfig.Password,
		migrationConfig.PostgresConfig.Host,
		migrationConfig.PostgresConfig.Port,
		migrationConfig.PostgresConfig.DbName)

	// migration argument
	migrationType := flag.String("type", "no-type", "type your migration")
	migrationVersion := flag.Int("version", 0, "version your migration")
	flag.Parse()

	if *migrationType == Up {
		migrateUp(migrationConfig.MigrationPath, dsn)
	} else if *migrationType == Down {
		migrateDown(migrationConfig.MigrationPath, dsn, *migrationVersion)
	} else if *migrationType == Force {
		migrateForce(migrationConfig.MigrationPath, dsn, *migrationVersion)
	} else {
		log.Println("use arguments to run the migration you need")
	}
}

// migrateUp is
func migrateUp(migrationPath, dbConn string) {
	m, err := migrate.New(migrationPath, dbConn)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = m.Up()
	if err != nil && errors.Is(err, migrate.ErrNoChange) {
		log.Println(err)
		return
	}

	log.Println("migrations up completed successfully")
}

// migrateDown is
func migrateDown(migrationPath, dbConn string, migrationVersion int) {
	if migrationVersion == 0 {
		log.Println("the arguments version given are incorrect for migration")
		return
	}
	m, err := migrate.New(migrationPath, dbConn)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = m.Migrate(uint(migrationVersion))
	if err != nil && errors.Is(err, migrate.ErrNoChange) {
		log.Println(err)
		return
	}

	log.Println("migrations down completed successfully")
}

// migrateForce is
func migrateForce(migrationPath, dbConn string, migrationVersion int) {
	if migrationVersion == 0 {
		log.Println("the arguments version given are incorrect for migrations")
		return
	}
	m, err := migrate.New(migrationPath, dbConn)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = m.Force(migrationVersion)
	if err != nil && errors.Is(err, migrate.ErrNoChange) {
		log.Println(err)
		return
	}

	log.Println("migrations force completed successfully")
}
