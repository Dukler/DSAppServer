package controllers

import (
	"DSAppServer/api"
	"DSAppServer/repositories/users"
	"DSAppServer/utils"
	"database/sql"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

var CreateUser = func(w http.ResponseWriter, r *http.Request) {
	usr, err := api.NewUser(w, r)
	usrRepo, _ := users.NewUserRepo()
	_, err = usrRepo.Create(usr)

	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	//api.Login(w,r)
	usr := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	email := usr["email"].(string)
	pwd := usr["password"].(string)
	// Get the existing entry present in the database for the given username or email
	storedCreds := &api.User{}

	usrRepo, _ := users.NewUserRepo()
	storedCreds, err = usrRepo.GetByEmail(email)

	if err != nil {
		// If an entry with the username does not exist, send an "Unauthorized"(401) status
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// If the error is of any other type, send a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Compare the stored hashed password, with the hashed version of the password that was received
	if err = bcrypt.CompareHashAndPassword([]byte(storedCreds.Password), []byte(pwd)); err != nil {
		// If the two passwords don't match, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// If we reach this point, that means the users password was correct, and that they are authorized
	// The default 200 status is sent
	resp := make(map[string]interface{})
	resp["email"] = storedCreds.Email
	resp["token"] = storedCreds.Token.String
	w.WriteHeader(http.StatusOK)
	utils.Respond(w, resp)
}
