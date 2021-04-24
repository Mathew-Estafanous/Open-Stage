package main

import (
	"context"
	"database/sql"
	"github.com/Mathew-Estafanous/Open-Stage/handler"
	"github.com/Mathew-Estafanous/Open-Stage/infrastructure/postgres"
	"github.com/Mathew-Estafanous/Open-Stage/service"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
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
}

func connectToDB() *sql.DB {
	dbUrl := os.Getenv("DATABASE_URL")
	sslDisabled := os.Getenv("SSL_DISABLED")
	if strings.ToLower(sslDisabled) == "true" {
		dbUrl += "?sslmode=disable"
	}

	db, err := sql.Open("postgres", dbUrl)
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
		ReadTimeout:  time.Second * 25,
		WriteTimeout: time.Second * 25,
	}
}
