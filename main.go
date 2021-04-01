package main

import (
	"database/sql"
	"github.com/Mathew-Estafanous/Open-Stage/database"
	"github.com/Mathew-Estafanous/Open-Stage/handler"
	"github.com/Mathew-Estafanous/Open-Stage/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const port string = ":8080"

func main() {
	db, err := sql.Open("mysql", "admin:qa-password@tcp(127.0.0.1:3306)/qa_platform")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Could not make a connection with the database.\n%v", err)
	}
	defer db.Close()

	roomStore := database.NewMySQLRoomStore(db)
	roomService := service.NewRoomService(roomStore)
	roomHandler := handler.NewRoomHandler(roomService)

	r := mux.NewRouter()
	roomHandler.HandleRoutes(r)

	log.Printf("Open-Stage starting on port %v", port)
	log.Fatal(http.ListenAndServe(port, r))
}
