package main

import (
	"context"
	"database/sql"
	"github.com/Mathew-Estafanous/Open-Stage/handler"
	"github.com/Mathew-Estafanous/Open-Stage/infrastructure/postgres"
	"github.com/Mathew-Estafanous/Open-Stage/infrastructure/redis"
	"github.com/Mathew-Estafanous/Open-Stage/middleware"
	"github.com/Mathew-Estafanous/Open-Stage/service"
	middle "github.com/go-openapi/runtime/middleware"
	red "github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	PROFILE = os.Getenv("PROFILE")
)

func main() {
	db := connectToDB()
	client := connectToRedis()

	redisCache := redis.NewMemoryCache(client)

	rStore := postgres.NewRoomStore(db)
	rService := service.NewRoomService(rStore)
	roomHandler := handler.NewRoomHandler(rService)

	qStore := postgres.NewQuestionStore(db)
	qService := service.NewQuestionService(qStore, rService)
	questionHandler := handler.NewQuestionHandler(qService)

	aStore := postgres.NewAccountStore(db)
	aService := service.NewAccountService(aStore)
	authService := service.NewAuthService(aStore, redisCache)
	accountHandler := handler.NewAccountHandler(aService, authService)

	router := mux.NewRouter()
	configureDocsRoute(router)

	apiRouter := router.PathPrefix("/v1").Subrouter()
	securedRouter := apiRouter.PathPrefix("/").Subrouter()
	securedRouter.Use(middleware.Auth(redisCache))

	roomHandler.Route(apiRouter, securedRouter)
	accountHandler.Route(apiRouter, securedRouter)
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

func connectToRedis() *red.Client {
	redisUrl := os.Getenv("REDIS_URL")
	redisOpt, err := red.ParseURL(redisUrl)
	if err != nil {
		log.Fatal(err.Error())
	}
	client := red.NewClient(redisOpt)

	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Fatal(err.Error())
	}
	return client
}

func connectToDB() *sql.DB {
	dbUrl := os.Getenv("DATABASE_URL")
	// If not a prod, so SSL is not required and cane be disabled.
	if PROFILE != "prod" {
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

func portByProfile() string {
	//If 'prod' profile, then use the assigned PORT env.
	if PROFILE == "prod" {
		return ":" + os.Getenv("PORT")
	}
	//Not 'prod', so use default 8080 port.
	return ":8080"
}
