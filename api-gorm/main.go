package main

import (
	"encoding/json"
	"fmt"
	"gorm-go/config"
	"net/http"

	"github.com/gorilla/mux"
)

type Contact struct {
	Id    uint
	Name  string
	Phone string
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/contact", GetAllContact).Methods("GET")
	router.HandleFunc("/contact/{id}", GetContactById).Methods("GET")
	router.HandleFunc("/contact", AddNewContact).Methods("POST")
	router.HandleFunc("/contact/{id}", EditContact).Methods("PUT")
	router.HandleFunc("/contact/{id}", DeleteContact).Methods("DELETE")

	fmt.Println("localhost:8080")
	http.ListenAndServe(":8080", router)
}

func GetAllContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := config.ConnectDB()
	if err != nil {
		http.Error(w, "Tidak dapat terconnect dengan database", http.StatusBadRequest)
		return
	}

	var read []Contact
	if err = db.Find(&read).Error; err != nil {
		http.Error(w, "Data contact tidak ditemukan", http.StatusBadRequest)
		return
	}

	result, err := json.Marshal(read)
	if err != nil {
		panic(err)
	}

	w.Write(result)
	http.Error(w, "", http.StatusBadRequest)
}

func GetContactById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := config.ConnectDB()
	if err != nil {
		http.Error(w, "Tidak dapat terconnect dengan database", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	readById := &Contact{}
	if err = db.First(&readById, "id = ?", id).Error; err != nil {
		http.Error(w, "Data tidak berhasil ditemukan", http.StatusBadRequest)
		return
	}

	result, err := json.Marshal(readById)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	w.Write(result)
	http.Error(w, "", http.StatusBadRequest)
}

func AddNewContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var contact Contact
	err := json.NewDecoder(r.Body).Decode(&contact)

	db, err := config.ConnectDB()
	if err != nil {
		http.Error(w, "Tidak dapat terconnect dengan database", http.StatusBadRequest)
		return
	}

	if err = db.Create(&contact).Error; err != nil {
		http.Error(w, "Data tidak dapat ditambahkan", http.StatusBadRequest)
		return
	}

	w.Write([]byte("Data berhasil ditambahkan"))
	http.Error(w, "", http.StatusBadRequest)
}

func EditContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var contact Contact
	err := json.NewDecoder(r.Body).Decode(&contact)

	vars := mux.Vars(r)
	id := vars["id"]

	db, err := config.ConnectDB()
	if err != nil {
		http.Error(w, "Tidak dapat terconnect dengan database", http.StatusBadRequest)
		return
	}

	if db.Where("id = ?", id).Updates(&contact).RowsAffected == 0 {
		http.Error(w, "Data tidak bisa di ubah", http.StatusBadRequest)
		return
	}

	w.Write([]byte("Data berhasil di ubah"))
	http.Error(w, "", http.StatusBadRequest)
}

func DeleteContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	db, err := config.ConnectDB()
	if err != nil {
		http.Error(w, "Tidak dapat terconnect dengan database", http.StatusBadRequest)
		return
	}

	var contact Contact
	if db.Delete(&contact, "id = ?", id).RowsAffected == 0 {
		http.Error(w, "Data tidak bisa dihapus", http.StatusBadRequest)
		return
	}

	w.Write([]byte("Data berhasil dihapus"))
	http.Error(w, "", http.StatusBadRequest)
}
