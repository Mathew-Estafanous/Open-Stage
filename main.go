package main

import (
	"context"
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
	"os/signal"
	"syscall"
	"time"
)

const port string = ":8080"

func main() {
	db := connectToDB()

	rStore := postgres.NewRoomStore(db)
	rService := service.NewRoomService(rStore)
	roomHandler := handler.NewRoomHandler(rService)

	qStore := postgres.NewQuestionStore(db)
	qService := service.NewQuestionService(qStore, rService)
	questionHandler := handler.NewQuestionHandler(qService)

	r := mux.NewRouter().PathPrefix("/v1").Subrouter()
	roomHandler.Route(r)
	questionHandler.Route(r)

	log.Printf("Open-Stage starting on port %v", port)
	server := configureServer(r)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGTERM)
		<-c

		log.Println("Shutting down server..")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
		defer cancel()

		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
		log.Println("Database connection closed.")

		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("Error during server shutdown: %+v", err)
		}
		log.Println("Server successfully shutdown.")
	}()

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
	log.Println("Server started")
}

func connectToDB() *sql.DB {
	port := os.Getenv("DATABASE_PORT")
	name := os.Getenv("DATABASE_NAME")
	address := os.Getenv("DATABASE_ADDRESS")
	user := os.Getenv("DATABASE_USERNAME")
	pass := os.Getenv("DATABASE_PASSWORD")

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", address, port, user, pass, name)
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

func configureServer(r http.Handler) *http.Server {
	return &http.Server{
		Addr: port,
		Handler: r,
		ReadTimeout:  time.Second * 20,
		WriteTimeout: time.Second * 20,
	}
}
