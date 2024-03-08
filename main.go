package main

import (
	"UTS/controller"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/rooms", controller.ShowAllRooms).Methods("GET")
	router.HandleFunc("/room", controller.ShowDetailRoom).Methods("GET")
	router.HandleFunc("/participants", controller.JoinRoom).Methods("POST")
	router.HandleFunc("/participants", controller.LeaveRoom).Methods("DELETE")

	http.Handle("/", router)
	fmt.Println("Connected to port 8080")
	log.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
