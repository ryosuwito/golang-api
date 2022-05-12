package handlers

import (
	"api-mux/structs"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var dbProducts structs.Products

	DB, _ := gorm.Open("mysql", "root:@/db_nasabah?charset=utf8&parseTime=True&loc=Local")

	json.Unmarshal(payloads, &dbProducts)
	DB.Create(&dbProducts)

	res := structs.Result{Code: 200, Data: dbProducts, Message: "Berhasil Menambahkan Product Baru"}

	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetProductsLimit(w http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	if limit < 1 {
		limit = 10
	}

	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	if limit < 1 {
		limit = 0
	}

	dbProducts := []structs.Products{}

	DB, _ := gorm.Open("mysql", "root:@/db_nasabah?charset=utf8&parseTime=True&loc=Local")

	DB.Limit(limit).Offset(offset).Find(&dbProducts)

	res := structs.Result{Code: 200, Data: dbProducts, Message: "User has successfully retrieve"}
	resuts, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resuts)
}

func GetProductId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	dbProducts := structs.Products{}
	DB, _ := gorm.Open("mysql", "root:@/db_nasabah?charset=utf8&parseTime=True&loc=Local")

	DB.First(&dbProducts, id)

	res := structs.Result{Code: 200, Data: dbProducts, Message: "Product Ditemukan"}

	result, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func UpdateProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	payloads, _ := ioutil.ReadAll(r.Body)

	var dbProducts structs.Products
	DB, _ := gorm.Open("mysql", "root:@/db_nasabah?charset=utf8&parseTime=True&loc=Local")

	DB.First(&dbProducts, id)

	json.Unmarshal(payloads, &dbProducts)

	DB.Model(&dbProducts).Update(dbProducts)

	res := structs.Result{Code: 200, Data: dbProducts, Message: "Berhasil Update Data Product"}

	result, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func DeleteProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var dbProducts structs.Products

	DB, _ := gorm.Open("mysql", "root:@/db_nasabah?charset=utf8&parseTime=True&loc=Local")

	DB.First(&dbProducts, id)
	DB.Delete(&dbProducts)

	res := structs.Result{Code: 200, Data: dbProducts, Message: "Berhasil Menghapus Product"}

	result, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
