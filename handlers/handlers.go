package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerReq() {
	log.Println("Start development server localhost:9000")

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", HomePage)
	// -- USERS--

	//Tambah User Baru
	r.HandleFunc("/user", CreateUsers).Methods("OPTIONS", "POST")

	//Find User Berdasarkan Limit
	r.HandleFunc("/users", GetUsersLimit).Methods("OPTIONS", "GET")

	//Find User Berdasarkan Id
	r.HandleFunc("/user/{id}", GetUserId).Methods("OPTIONS", "GET")

	//Update User Berdasarkan Id
	r.HandleFunc("/user/{id}", UpdateUserById).Methods("OPTIONS", "PUT")

	//Hapus User Berdasarkan Id
	r.HandleFunc("/user/{id}", DeleteUserById).Methods("OPTIONS", "DELETE")

	//Login User Berdasarkan Id
	r.HandleFunc("/login", LoginUser).Methods("OPTIONS", "POST")

	// -- USERS --

	// -- PRODUCT --

	//Tambah Product Baru
	r.HandleFunc("/product", CreateProduct).Methods("OPTIONS", "POST")

	//Find Product Berdasarkan Limit
	r.HandleFunc("/products", GetProductsLimit).Methods("OPTIONS", "GET")

	//Find Product Berdasarkan Id
	r.HandleFunc("/product/{id}", GetProductId).Methods("OPTIONS", "GET")

	//Update Product Berdasarkan Id
	r.HandleFunc("/product/{id}", UpdateProductById).Methods("OPTIONS", "PUT")

	//Hapus Product Berdasarkan Id
	r.HandleFunc("/product/{id}", DeleteProductById).Methods("OPTIONS", "DELETE")

	// -- PRODUCT --

	log.Fatal(http.ListenAndServe(":9000", r))
}
