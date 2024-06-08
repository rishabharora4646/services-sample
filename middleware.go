package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/mitchellh/mapstructure"
)

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("x-auth-token")
		if authorizationHeader != "" {
			token, error := jwt.Parse(authorizationHeader, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("there was an error")
				}
				return []byte(JWT_SECRET), nil
			})
			if error != nil {
				resp := &errorResponse{
					Message: "Invalid Auth Token.",
				}
				b := new(bytes.Buffer)
				json.NewEncoder(b).Encode(resp)
				w.Header().Set("Content-Type", "application/json")
				http.Error(w, b.String(), http.StatusUnauthorized)
				return
			}
			if token.Valid {
				context.Set(req, "decoded", token.Claims)
				next(w, req)
			} else {
				resp := &errorResponse{
					Message: "Invalid Auth Token.",
				}
				b := new(bytes.Buffer)
				json.NewEncoder(b).Encode(resp)
				w.Header().Set("Content-Type", "application/json")
				http.Error(w, b.String(), http.StatusUnauthorized)
			}
		} else {
			resp := &errorResponse{
				Message: "Auth Token Header empty.",
			}
			b := new(bytes.Buffer)
			json.NewEncoder(b).Encode(resp)
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, b.String(), http.StatusUnauthorized)
		}
	})
}

func corsMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			//whitelist cors origin for security.
			if origin == "http://localhost" {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			}
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE, HEAD")
			w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Requested-With, Accept, Accept-Language, Accept-Encoding, Content-Language, x-auth-token, FIREBASE-AUTH-TOKEN")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func getJwtClaims(r *http.Request) jwtClaims {
	decoded := context.Get(r, "decoded")
	var claims jwtClaims
	mapstructure.Decode(decoded.(jwt.MapClaims), &claims)
	return claims
}

// helpers
type jwtClaims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}
