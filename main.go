package main

import (
	"database/sql"
	"fmt"
	"github.com/Mathew-Estafanous/Open-Stage/database"
	"github.com/Mathew-Estafanous/Open-Stage/handler"
	"github.com/Mathew-Estafanous/Open-Stage/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

const port string = ":8080"

func main() {
	db := connectToDB()
	defer db.Close()

	roomStore := database.NewMySQLRoomStore(db)
	roomService := service.NewRoomService(roomStore)
	roomHandler := handler.NewRoomHandler(roomService)

	r := mux.NewRouter()
	roomHandler.Route(r)

	log.Printf("Open-Stage starting on port %v", port)
	log.Fatal(http.ListenAndServe(port, r))
}

func connectToDB() *sql.DB {
	dbPort := os.Getenv("DATABASE_PORT")
	dbName := os.Getenv("DATABASE_NAME")
	dbAddress := os.Getenv("DATABASE_ADDRESS")
	user := os.Getenv("DATABASE_USERNAME")
	pass := os.Getenv("DATABASE_PASSWORD")

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", user, pass, dbAddress, dbPort, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Could not make a connection with the database.\n%v", err)
	}
	return db
}
