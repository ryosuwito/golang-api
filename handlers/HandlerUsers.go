package handlers

import (
	"api-mux/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

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
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	if limit < 1 {
		limit = 10
	}

	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	if limit < 1 {
		limit = 0
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

		if !dbUsers.Status {
			DB.Model(&dbUsers).Update(&dbUsers)
			DB.Model(&dbUsers).Updates(map[string]interface{}{"status": false})
		}
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
		if !dbUsers.Status {
			DB.Model(&dbUsers).Update(&dbUsers)
			DB.Model(&dbUsers).Updates(map[string]interface{}{"status": false})
		}
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
	var userLogin structs.UsersLogin
	DB, _ := gorm.Open("mysql", "root:@/db_nasabah?charset=utf8&parseTime=True&loc=Local")

	json.Unmarshal(payloads, &dbUsers)
	DB.Where("email = ? AND password >= ?", &dbUsers.Email, &dbUsers.Password).Find(&dbUsers)

	userLogin.ID = dbUsers.ID
	userLogin.Email = dbUsers.Email
	userLogin.Password = dbUsers.Password
	res := structs.Result{Code: 200, Data: userLogin, Message: "Berhasil Login"}

	result, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
