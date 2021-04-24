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
	"syscall"
	"time"
)

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

	port := PortByProfile()
	log.Printf("Open-Stage starting on port %v", port)
	server := configureServer(r, port)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGTERM)
		<-c

		log.Println("Shutting down server..")
		ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
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
	dbUrl := DbUrlByProfile()
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

func configureServer(r http.Handler, port string) *http.Server {
	return &http.Server{
		Addr: port,
		Handler: r,
		ReadTimeout:  25 * time.Second,
		WriteTimeout: 25 * time.Second,
	}
}
