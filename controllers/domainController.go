package controllers

import (
	"DSLoginServer/api"
	"DSLoginServer/repositories/domains"
	"DSLoginServer/utils"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"io"
	"log"
	"net/http"
	"os"
)

var DomainApp = func(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	body := decodeBody(r.Body)
	token := body["token"].(string)
	domStr := body["domain"].(string)
	_,valid := extractClaims(token)
	domRepo, _ :=  domains.NewDomainRepo()
	dom := &api.Domain{}
	var err error

	dom, err = domRepo.GetByName(domStr)
	if err != nil {
		panic(err)
	}

	if valid {
		data["loggedIn"] = true
	}

	data["appName"] = dom.AppName
	utils.Respond(w,data)
}

func decodeBody(body io.ReadCloser) map[string]interface{}{
	decoder := json.NewDecoder(body)
	var b map[string]interface{}
	err := decoder.Decode(&b)
	if err != nil {
		panic(err)
	}
	return b
}

func extractClaims(tokenStr string) (jwt.MapClaims, bool) {
	hmacSecretString := os.Getenv("token_password")
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}