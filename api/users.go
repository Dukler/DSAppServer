package api

import (
	"database/sql"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"time"
)


/*
JWT claims struct
*/
type Token struct {
	Email string
	jwt.StandardClaims
}

//a struct to rep users
type User struct {
	ID []uint8 `db:"id"`
	Email string `db:"email"`
	Username string `db:"username"`
	Password string `db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	Token sql.NullString `db:"token"`
}

func NewUser(w http.ResponseWriter, r *http.Request) (*User, error){
	// Parse and decode the request body into a new `User` instance
	//db := dbh.GetDB()
	usr := &User{}
	err := json.NewDecoder(r.Body).Decode(usr)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return nil, err
	}

	tk := &Token{Email: usr.Email}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	log.Print(os.Getenv("token_password"))
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	usr.Token = sql.NullString{String:tokenString, Valid:true}

	// Salt and hash the password using the bcrypt algorithm
	// The second argument is the cost of hashing, which we arbitrarily set as 8 (this value can be more or less, depending on the computing power you wish to utilize)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usr.Password), 8)
	usr.Password = string(hashedPassword)

	return usr, err
}
