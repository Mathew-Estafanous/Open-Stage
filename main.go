package main

import (
	"database/sql"
	"fmt"
	"github.com/Mathew-Estafanous/Open-Stage/handler"
	"github.com/Mathew-Estafanous/Open-Stage/infrastructure/postgres"
	"github.com/Mathew-Estafanous/Open-Stage/service"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

const port string = ":8080"

func main() {
	db := connectToDB()
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	rStore := postgres.NewRoomStore(db)
	rService := service.NewRoomService(rStore)
	roomHandler := handler.NewRoomHandler(rService)

	qStore := postgres.NewQuestionStore(db)
	qService := service.NewQuestionService(qStore, rService)
	questionHandler := handler.NewQuestionHandler(qService)

	r := mux.NewRouter()
	roomHandler.Route(r)
	questionHandler.Route(r)

	log.Printf("Open-Stage starting on port %v", port)
	log.Fatal(http.ListenAndServe(port, r))
}

func connectToDB() *sql.DB {
	dbPort := os.Getenv("DATABASE_PORT")
	dbName := os.Getenv("DATABASE_NAME")
	dbAddress := os.Getenv("DATABASE_ADDRESS")
	user := os.Getenv("DATABASE_USERNAME")
	pass := os.Getenv("DATABASE_PASSWORD")

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", dbAddress, dbPort, user, pass, dbName)
	log.Println(dsn)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Could not make a connection with the database.\n%v", err)
	}
	return db
}
