package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func servicesListHandler(w http.ResponseWriter, r *http.Request) {

	claims := getJwtClaims(r)

	params := mux.Vars(r)
	orgId := params["orgId"]

	if checkUserCanAccessOrg(orgId, claims.UserId) {
		v := r.URL.Query()
		offset := v.Get("offset")
		name := v.Get("name")

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
		var selDB *sql.Rows

		if len(name) > 0 {
			selDB, err = db.Query("SELECT services.service_id,services.service_name,services.service_description,services.created_at,services.updated_at,COUNT(services_versions.version_id) FROM services,services_versions WHERE services.service_id=services_versions.service_id AND services.org_id=? AND services.service_name LIKE ? GROUP BY services.service_id LIMIT ?,25", orgId, "%"+name+"%", offset)
		} else {
			selDB, err = db.Query("SELECT services.service_id,services.service_name,services.service_description,services.created_at,services.updated_at,COUNT(services_versions.version_id) FROM services,services_versions WHERE services.service_id=services_versions.service_id AND services.org_id=? GROUP BY services.service_id LIMIT ?,25", orgId, offset)
		}
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
		resp := []serviceListResponse{}
		for selDB.Next() {
			data := serviceListResponse{}
			err = selDB.Scan(&data.ServiceId, &data.ServiceName, &data.ServiceDescription, &data.CreatedAt, &data.UpdatedAt, &data.VersionCount)
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
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	} else {
		resp := &errorResponse{
			Message: "User does not have access to this org.",
		}
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(resp)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, b.String(), http.StatusForbidden)
		return
	}
}

func servicesInfoHandler(w http.ResponseWriter, r *http.Request) {

	claims := getJwtClaims(r)

	params := mux.Vars(r)
	serviceId := params["serviceId"]

	if checkUserCanAccessService(serviceId, claims.UserId) {

		selDB, err := db.Query("SELECT services.service_id,services.service_name,services.service_description,services.created_at,services.updated_at,COUNT(services_versions.version_id) FROM services,services_versions WHERE services.service_id=services_versions.service_id AND services.service_id=?", serviceId)
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
		resp := serviceInfoResponse{}
		for selDB.Next() {
			err = selDB.Scan(&resp.ServiceId, &resp.ServiceName, &resp.ServiceDescription, &resp.CreatedAt, &resp.UpdatedAt, &resp.VersionCount)
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

		//get versions for service
		selDB, err = db.Query("SELECT version_id,version_name,service_host,service_port,is_active,created_at,updated_at FROM services_versions WHERE service_id=?", serviceId)
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
		resp.Versions = []serviceVersionInfo{}
		for selDB.Next() {
			data := serviceVersionInfo{}
			var isActive int
			err = selDB.Scan(&data.VersionId, &data.VersionName, &data.ServiceHost, &data.ServicePort, &isActive, &data.CreatedAt, &data.UpdatedAt)
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
			if isActive == 1 {
				data.Is_active = true
			}
			resp.Versions = append(resp.Versions, data)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	} else {
		resp := &errorResponse{
			Message: "User does not have access to this org.",
		}
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(resp)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, b.String(), http.StatusForbidden)
		return
	}
}

// helpers
func checkUserCanAccessOrg(org_id, user_id string) (state bool) {
	err := db.QueryRow("SELECT exists (SELECT org_id FROM org_users WHERE org_id=? AND user_id=?)", org_id, user_id).Scan(&state)
	if err != nil && err != sql.ErrNoRows {
		log.Println("error checking if row exists ", err)
	}
	return state
}

func checkUserCanAccessService(service_id, user_id string) (state bool) {
	err := db.QueryRow("SELECT exists (SELECT services.org_id FROM services,org_users WHERE services.org_id=org_users.org_id AND services.service_id=? AND org_users.user_id=?)", service_id, user_id).Scan(&state)
	if err != nil && err != sql.ErrNoRows {
		log.Println("error checking if row exists ", err)
	}
	return state
}

type serviceListResponse struct {
	ServiceId          string `json:"serviceId"`
	ServiceName        string `json:"serviceName"`
	ServiceDescription string `json:"serviceDescription"`
	CreatedAt          string `json:"createdAt"`
	UpdatedAt          string `json:"updatedAt"`
	VersionCount       int    `json:"versionCount"`
}

type serviceInfoResponse struct {
	ServiceId          string               `json:"serviceId"`
	ServiceName        string               `json:"serviceName"`
	ServiceDescription string               `json:"serviceDescription"`
	CreatedAt          string               `json:"createdAt"`
	UpdatedAt          string               `json:"updatedAt"`
	VersionCount       int                  `json:"versionCount"`
	Versions           []serviceVersionInfo `json:"versions"`
}

type serviceVersionInfo struct {
	VersionId   string `json:"versionId"`
	VersionName string `json:"versionName"`
	ServiceHost string `json:"serviceHost"`
	ServicePort int    `json:"servicePort"`
	Is_active   bool   `json:"isActive"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}
