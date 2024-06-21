package db

// import (
// 	"log"
// 	"os"
// 	"time"

// 	"github.com/jackc/pgx/v5"
// 	"github.com/jackc/pgx/v5/stdlib"
// 	"github.com/jmoiron/sqlx"
// 	"github.com/joho/godotenv"
// )

// var DB *sqlx.DB

// func Init() {
// 	// Load environment variables from .env file
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	// Get PostgreSQL connection URI from environment variable
// 	pgURI := os.Getenv("PG_URI")
// 	if pgURI == "" {
// 		log.Fatal("PG_URI environment variable not set")
// 	}

// 	// Parse PG_URI to pgx connection config
// 	connConfig, err := pgx.ParseConfig(pgURI)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Open stdlib database connection
// 	pgxStdlibDB := stdlib.OpenDB(*connConfig)

// 	// Create sqlx database object
// 	DB = sqlx.NewDb(pgxStdlibDB, "pgx")
// 	DB.SetMaxIdleConns(4)
// 	DB.SetMaxOpenConns(8)
// 	DB.SetConnMaxLifetime(time.Duration(30) * time.Minute)

// 	// Create database tables if they do not exist
// 	createTables()

// 	log.Println("DB connected!")
// }

// func createTables() {
// 	createEventsTable := `
// 		CREATE TABLE "accounts" (
// 			"id" bigserial PRIMARY KEY,
// 			"owner" varchar NOT NULL,
// 			"balance" numeric NOT NULL,
// 			"currency" varchar NOT NULL,
// 			"created_at" timestamp NOT NULL DEFAULT (now())
// 		);

// 		CREATE TABLE "entries" (
// 			"id" bigserial PRIMARY KEY,
// 			"account_id" bigint NOT NULL,
// 			"amount" numeric NOT NULL,
// 			"created_at" timestamp NOT NULL DEFAULT (now())
// 		);

// 		CREATE TABLE "transfers" (
// 			"id" bigserial PRIMARY KEY,
// 			"from_account_id" bigint,
// 			"to_account_id" bigint,
// 			"amount" numeric NOT NULL,
// 			"created_at" timestamp NOT NULL DEFAULT (now())
// 		);

// 		CREATE INDEX ON "accounts" ("owner");

// 		CREATE INDEX ON "entries" ("account_id");

// 		CREATE INDEX ON "transfers" ("from_account_id");

// 		CREATE INDEX ON "transfers" ("to_account_id");

// 		CREATE INDEX ON "transfers" ("from_account_id", "to_account_id");

// 		COMMENT ON COLUMN "entries"."amount" IS 'can be -ve/+ve';

// 		COMMENT ON COLUMN "transfers"."amount" IS 'must be +ve';

// 		ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

// 		ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

// 		ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");
// 	`

// 	_, err := DB.Exec(createEventsTable)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
