package main

import (
	"api-mux/connections"
	"api-mux/handlers"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	connections.Connect()
	handlers.HandlerReq()
}

/* func Connection() {

	DB, Err := gorm.Open("mysql", "root:@/db_nasabah?charset=utf8&parseTime=True&loc=Local")
	if Err != nil {
		fmt.Println("Gagal Koneksi", Err)
	} else {
		fmt.Println("Berhasil Koneksi")
	}
	DB.AutoMigrate(&Product{})
}

func CreateProduct() {
} */
