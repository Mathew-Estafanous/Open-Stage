package main

import (
	"context"
	"database/sql"
	"github.com/Mathew-Estafanous/Open-Stage/handler"
	"github.com/Mathew-Estafanous/Open-Stage/infrastructure/postgres"
	"github.com/Mathew-Estafanous/Open-Stage/middleware"
	"github.com/Mathew-Estafanous/Open-Stage/service"
	middle "github.com/go-openapi/runtime/middleware"
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

func main() {
	db := connectToDB()

	rStore := postgres.NewRoomStore(db)
	rService := service.NewRoomService(rStore)
	roomHandler := handler.NewRoomHandler(rService)

	qStore := postgres.NewQuestionStore(db)
	qService := service.NewQuestionService(qStore, rService)
	questionHandler := handler.NewQuestionHandler(qService)

	router := mux.NewRouter()
	configureDocsRoute(router)

	apiRouter := router.PathPrefix("/v1").Subrouter()
	roomHandler.Route(apiRouter)
	questionHandler.Route(apiRouter)

	port := portByProfile()
	log.Printf("Open-Stage starting on port %v", port)
	server := configureServer(middleware.CORS()(router), port)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGTERM)
		<-c

		log.Println("Shutting down server..")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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
	dbUrl := dbUrlByProfile()
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

func configureDocsRoute(router *mux.Router) {
	opts := middle.RedocOpts{
		SpecURL: "/docs/swagger.yaml",
		Title:   "Open-Stage API Docs",
	}
	doc := middle.Redoc(opts, nil)
	router.Handle("/docs", doc)
	router.Handle("/docs/swagger.yaml", http.FileServer(http.Dir("./")))
}

func configureServer(r http.Handler, port string) *http.Server {
	return &http.Server{
		Addr:         port,
		Handler:      middleware.EnforceSSL(r),
		ReadTimeout:  25 * time.Second,
		WriteTimeout: 25 * time.Second,
	}
}

func portByProfile() string {
	//If 'prod' profile, then use the assigned PORT env.
	if os.Getenv("PROFILE") == "prod" {
		return ":" + os.Getenv("PORT")
	}
	//Not 'prod', so use default 8080 port.
	return ":8080"
}

func dbUrlByProfile() string {
	dbUrl := os.Getenv("DATABASE_URL")
	//If the 'prod' profile, return the given database_url
	if os.Getenv("PROFILE") == "prod" {
		return dbUrl
	}

	//Not a prod, so SSL is not required and cane be disabled.
	dbUrl += "?sslmode=disable"
	//Check if profile is a container. If so, replace address with container address.
	if os.Getenv("PROFILE") == "ctr" {
		ctrAddr := os.Getenv("CONTAINER_ADDRESS")
		return strings.Replace(dbUrl, "[address]", ctrAddr, 1)
	}
	//If profile isn't any of the above, assume 'dev' and return altered dbUrl.
	return dbUrl
}
