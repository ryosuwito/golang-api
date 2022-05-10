package handlers

import (
	"api-mux/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wilkommen!")
}
func CreateUsers(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var dbUsers structs.Users

	DB, _ := gorm.Open("mysql", "root:@/db_nasabah?charset=utf8&parseTime=True&loc=Local")

	json.Unmarshal(payloads, &dbUsers)
	if dbUsers.Role == "0" {
		dbUsers.Role = "user"
		DB.Create(&dbUsers)

		res := structs.Result{Code: 200, Data: dbUsers, Message: "Berhasil Menambahkan User Baru"}

		result, err := json.Marshal(res)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	} else if dbUsers.Role == "1" {
		dbUsers.Role = "admin"
		DB.Create(&dbUsers)

		res := structs.Result{Code: 200, Data: dbUsers, Message: "Berhasil Menambahkan User Baru"}

		result, err := json.Marshal(res)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	} else {
		res := structs.Result{Code: 200, Data: dbUsers, Message: "Gagal Memasukan User baru Karena Pengisian Role salah. Isi Role 0 atau 1"}

		result, err := json.Marshal(res)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}
}

func GetUsersLimit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var limit interface{}
	var offset interface{}

	limit = vars["limit"]
	offset = vars["offset"]

	if limit == "" {
		limit = 10
	}

	if offset == "" {
		offset = 0
	}

	dbUsers := []structs.Users{}

	DB, _ := gorm.Open("mysql", "root:@/db_nasabah?charset=utf8&parseTime=True&loc=Local")

	DB.Limit(limit).Offset(offset).Find(&dbUsers)

	res := structs.Result{Code: 200, Data: dbUsers, Message: "User has successfully retrieve"}
	resuts, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resuts)
}

func GetUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	dbUsers := structs.Users{}
	DB, _ := gorm.Open("mysql", "root:@/db_nasabah?charset=utf8&parseTime=True&loc=Local")

	DB.First(&dbUsers, id)

	res := structs.Result{Code: 200, Data: dbUsers, Message: "Users Ditemukan"}

	result, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func UpdateUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	payloads, _ := ioutil.ReadAll(r.Body)

	var dbUsers structs.Users
	DB, _ := gorm.Open("mysql", "root:@/db_nasabah?charset=utf8&parseTime=True&loc=Local")

	DB.First(&dbUsers, id)

	json.Unmarshal(payloads, &dbUsers)
	if dbUsers.Role == "0" {
		dbUsers.Role = "users"
		DB.Model(&dbUsers).Update(dbUsers)

		res := structs.Result{Code: 200, Data: dbUsers, Message: "Berhasil Update Data Users"}

		result, err := json.Marshal(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	} else if dbUsers.Role == "1" {
		dbUsers.Role = "admin"
		DB.Model(&dbUsers).Update(dbUsers)

		res := structs.Result{Code: 200, Data: dbUsers, Message: "Berhasil Update Data Users"}

		result, err := json.Marshal(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	} else {
		res := structs.Result{Code: 200, Data: dbUsers, Message: "Gagal Update User Karena Pengisian Role salah. Isi Role 0 atau 1"}

		result, err := json.Marshal(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}
}

func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var dbUsers structs.Users

	DB, _ := gorm.Open("mysql", "root:@/db_nasabah?charset=utf8&parseTime=True&loc=Local")

	DB.First(&dbUsers, id)
	DB.Delete(&dbUsers)

	res := structs.Result{Code: 200, Data: dbUsers, Message: "Berhasil Menghapus Users"}

	result, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var dbUsers structs.Users
	DB, _ := gorm.Open("mysql", "root:@/db_nasabah?charset=utf8&parseTime=True&loc=Local")

	json.Unmarshal(payloads, &dbUsers)
	DB.Where("email = ? AND password >= ?", "jinzhu", "22").Find(&dbUsers)

	res := structs.Result{Code: 200, Data: dbUsers, Message: "Berhasil Login"}

	result, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
