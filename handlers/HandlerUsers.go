package handlers

import (
	"api-mux/connections"
	"api-mux/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wilkommen!")
}

func CreateUsers(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)
	var dbUsers structs.Users

	res := structs.Result{Code: 500, Data: dbUsers, Message: "Unknown Error"}
	json.Unmarshal(payloads, &dbUsers)

	switch dbUsers.Role {
	case "0":
		dbUsers.Role = "user"
	case "1":
		dbUsers.Role = "admin"
	default:
		dbUsers.Role = "invalid"
		res.Code = 400
		res.Message = "Invalid User Role"
	}

	if dbUsers.Role != "invalid" {
		hashedPassword, err := HashPassword(dbUsers.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		dbUsers.Password = hashedPassword
		if err := connections.DB.Create(&dbUsers).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		res.Data = dbUsers
		res.Code = 200
		res.Message = "Add new user successfully"
	}

	result, _ := json.Marshal(res)
	ReturnResult(w, result)
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

	connections.DB.Limit(limit).Offset(offset).Find(&dbUsers)

	res := structs.Result{Code: 200, Data: dbUsers, Message: "User has successfully retrieve"}
	result, err := json.Marshal(res)
	ReturnCheckError(w, err)
	ReturnResult(w, result)
}

func GetUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	dbUsers := structs.Users{}

	connections.DB.First(&dbUsers, id)

	res := structs.Result{Code: 200, Data: dbUsers, Message: "Users Ditemukan"}

	result, err := json.Marshal(res)
	ReturnCheckError(w, err)
	ReturnResult(w, result)
}

func UpdateUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	payloads, _ := ioutil.ReadAll(r.Body)

	var dbUsers structs.Users

	res := structs.Result{Code: 200, Data: dbUsers, Message: "Unknown Error"}

	connections.DB.First(&dbUsers, id)

	json.Unmarshal(payloads, &dbUsers)

	switch dbUsers.Role {
	case "0":
		dbUsers.Role = "user"
	case "1":
		dbUsers.Role = "admin"
	default:
		dbUsers.Role = "invalid"
		res.Code = 400
		res.Message = "Invalid User Role"
	}

	if dbUsers.Role != "invalid" {
		if err := connections.DB.Model(&dbUsers).Update(&dbUsers).Error; err != nil {
			ReturnCheckError(w, err)
		}
		if !dbUsers.Status {
			connections.DB.Model(&dbUsers).Updates(map[string]interface{}{"status": false})
		}
		res.Code = 200
		res.Data = dbUsers
		res.Message = "Update user data successfully"
	}

	result, err := json.Marshal(res)
	ReturnCheckError(w, err)
	ReturnResult(w, result)

}

func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var dbUsers structs.Users

	connections.DB.First(&dbUsers, id)
	connections.DB.Delete(&dbUsers)

	res := structs.Result{Code: 200, Data: dbUsers, Message: "Berhasil Menghapus Users"}

	result, err := json.Marshal(res)
	ReturnCheckError(w, err)
	ReturnResult(w, result)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var dbUser structs.Users
	var userLogin structs.UsersLogin
	res := structs.Result{Code: 200, Data: userLogin, Message: "Gagal Login"}
	json.Unmarshal(payloads, &userLogin)
	connections.DB.Where("email = ?", &userLogin.Email).Find(&dbUser)

	if CheckPasswordHash(userLogin.Password, dbUser.Password) {
		res = structs.Result{Code: 200, Data: userLogin, Message: "Berhasil Login"}
	}

	result, err := json.Marshal(res)
	ReturnCheckError(w, err)
	ReturnResult(w, result)
}

func ReturnCheckError(w http.ResponseWriter, err error) {
	if err != nil {
		res := structs.Result{Code: http.StatusInternalServerError, Data: nil, Message: err.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		result, _ := json.Marshal(res)
		w.Write(result)
	}
}
func ReturnResult(w http.ResponseWriter, result []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
