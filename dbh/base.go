package dbh

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"path/filepath"
)

var db *sqlx.DB

type Test struct {
	//ID string `db:"id"`
	//Email  string `db:"email"`

	ID uint `sql:"id"`
	Email string `sql:"email"`
	Username string `sql:"username"`
	Password string `sql:"password"`
	//Token string `json:"token";sql:"-"`
	Token string `sql:"token"`
}

func InitDB() {
	path,_ := filepath.Abs("./utils/connection.env")
	//migrationsPath,_ := filepath.Abs("./migrations")
	e := godotenv.Load(path) //Load .env file
	if e != nil {
		fmt.Print(e)
	}
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	port := 5432

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, port, username, password, dbName)

	var err error
	db, err = sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//defer db.Close()


	driver, _ := postgres.WithInstance(db.DB, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"login", driver)

	if err != nil {
		log.Fatalf("migration failed... %v", err)
	}
	m.Steps(6)

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string
	fmt.Println(dbUri)
}

//returns a handle to the DB object
func GetDB() *sqlx.DB {
	return db
}
