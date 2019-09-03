package dbh

import (
	"errors"
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
	"time"
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

	ID uint `sql:"id"`
	Email string `sql:"email"`
	Username string `sql:"username"`
	Password string `sql:"password"`
	//Token string `json:"token";sql:"-"`
	Token string `sql:"token"`
}
func connectDB(psqlInfo string) bool{
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
	path,_ := filepath.Abs("./utils/connection.env")
	//migrationsPath,_ := filepath.Abs("./migrations")
	var connectionString string
	if os.Getenv("ENVIROMENT")=="LOCAL"{
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
	}else{
		connectionString = os.Getenv("DATABASE_URL")
	}



	test := connectDB(connectionString)
	reconnections := 5
	for i:=0; i < reconnections && !test; i++ {
		timer1 := time.NewTimer(5 * time.Second)
		<-timer1.C
		connectDB(connectionString)
	}
	var err error
	err = db.Ping()
	s := fmt.Sprintf("DB connection failed after %s tries", reconnections)
	if err != nil{
		dbErr := errors.New(s)
		panic(dbErr)
	}

	//defer db.Close()


	driver, _ := postgres.WithInstance(db.DB, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"login", driver)

	if err != nil {
		log.Fatalf("migration failed... %v", err)
	}
	err = m.Steps(7)
	if err != nil {
		panic(err)
	}


	fmt.Println(connectionString)
}


//returns a handle to the DB object
func GetDB() *sqlx.DB {
	return db
}
