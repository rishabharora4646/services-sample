package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func usersLoginHandler(w http.ResponseWriter, r *http.Request) {
	var p usersLoginRequest
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		resp := &errorResponse{
			Message: "Error Decoding Paylod.",
		}
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(resp)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, b.String(), http.StatusConflict)
		return
	}

	if p.Email != "" && p.Password != "" {
		var userId, hashedPassword string
		selDB, err := db.Query("SELECT user_id,password FROM users WHERE email=?", p.Email)
		if err != nil {
			log.Println(err.Error())
			resp := &errorResponse{
				Message: "Server Error. Try Again Later.",
			}
			b := new(bytes.Buffer)
			json.NewEncoder(b).Encode(resp)
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, b.String(), 500)
			return
		}
		for selDB.Next() {
			err = selDB.Scan(&userId, &hashedPassword)
			if err != nil {
				log.Println(err.Error())
				resp := &errorResponse{
					Message: "Server Error. Try Again Later.",
				}
				b := new(bytes.Buffer)
				json.NewEncoder(b).Encode(resp)
				w.Header().Set("Content-Type", "application/json")
				http.Error(w, b.String(), 500)
				return
			}
		}

		//comapre hash
		err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(p.Password))
		if err != nil {
			resp := &errorResponse{
				Message: "Invalid User Credentials.",
			}
			b := new(bytes.Buffer)
			json.NewEncoder(b).Encode(resp)
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, b.String(), http.StatusForbidden)
			return
		}

		//if no error create ans send jwt.
		claims := jwtClaims{
			userId,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * time.Duration(2160)).Unix(),
			},
		}
		jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, error := jwtToken.SignedString([]byte(JWT_SECRET))
		if error != nil {
			log.Println(error)
			resp := &errorResponse{
				Message: "Server Error. Try Again Later.",
			}
			b := new(bytes.Buffer)
			json.NewEncoder(b).Encode(resp)
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, b.String(), 500)
			return
		}
		resp := &usersLoginResponse{
			Token: tokenString,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	} else {
		resp := &errorResponse{
			Message: "Email and password are required.",
		}
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(resp)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, b.String(), http.StatusBadRequest)
		return
	}
}

// helpers
type usersLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type usersLoginResponse struct {
	Token string `json:"token"`
}
