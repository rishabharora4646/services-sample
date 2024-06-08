package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func orgsListHandler(w http.ResponseWriter, r *http.Request) {

	v := r.URL.Query()
	offset := v.Get("offset")

	_, err := strconv.Atoi(offset)

	if err != nil {
		log.Println(err.Error())
		resp := &errorResponse{
			Message: "Invalid Offset.",
		}
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(resp)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, b.String(), http.StatusBadRequest)
		return
	}

	claims := getJwtClaims(r)

	selDB, err := db.Query("SELECT orgs.org_id,orgs.org_name,orgs.created_at FROM orgs,org_users WHERE orgs.org_id=org_users.org_id AND org_users.user_id=? LIMIT ?,25", claims.UserId, offset)
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
	resp := []orgListResponse{}
	for selDB.Next() {
		data := orgListResponse{}
		err = selDB.Scan(&data.OrgId, &data.OrgName, &data.CreatedAt)
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
		resp = append(resp, data)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}

// helpers
type orgListResponse struct {
	OrgId     string `json:"orgId"`
	OrgName   string `json:"orgName"`
	CreatedAt string `json:"createdAt"`
}
