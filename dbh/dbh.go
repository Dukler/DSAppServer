package dbh

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

type Test struct {
	//ID string `db:"id"`
	//Email  string `db:"email"`

	ID       uint   `sql:"id"`
	Email    string `sql:"email"`
	Username string `sql:"username"`
	Password string `sql:"password"`
	//Token string `json:"token";sql:"-"`
	Token string `sql:"token"`
}

func connectDB(psqlInfo string) bool {
	var err error
	db, err = sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		//log.Fatal(err)
		fmt.Printf(ErrorColor, err)
		return false
	}
	return true
}

func InitDB() {
	path, _ := filepath.Abs("./utils/connection.env")
	//migrationsPath,_ := filepath.Abs("./migrations")
	var connectionString string
	env := os.Getenv("ENV")
	if env == "HEROKU" {
		connectionString = os.Getenv("DATABASE_URL")
	} else {
		e := godotenv.Load(path) //Load .env file
		if e != nil {
			fmt.Print(e)
		}
		username := os.Getenv("db_user")
		password := os.Getenv("db_pass")
		dbName := os.Getenv("db_name")
		dbHost := os.Getenv("db_host")
		port := 5432

		connectionString = fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			dbHost, port, username, password, dbName)
	}

	test := connectDB(connectionString)
	reconnections := 5
	for i := 0; i < reconnections && !test; i++ {
		timer1 := time.NewTimer(5 * time.Second)
		<-timer1.C
		connectDB(connectionString)
	}
	var err error
	err = db.Ping()
	if err != nil {
		s := fmt.Sprintf("DB connection failed after %s tries", reconnections)
		dbErr := errors.New(s)
		panic(dbErr)
	}

	//defer db.Close() no

	pgConfig := postgres.Config{}
	driver, err := postgres.WithInstance(db.DB, &pgConfig)
	if err != nil {
		log.Fatalf("could not start sql migration... %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		pgConfig.DatabaseName, driver)

	if err != nil {
		log.Fatalf("migration failed... %v", err)
	}
	//err = m.Down()
	//if err := m.Down(); err != nil && err != migrate.ErrNoChange {
	//	log.Fatalf("An error occurred while syncing the database.. %v", err)
	//}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("An error occurred while syncing the database.. %v", err)
	}

	
	// fmt.Println(connectionString) print connection string
}

//returns a handle to the DB object
func GetDB() *sqlx.DB {
	return db
}
