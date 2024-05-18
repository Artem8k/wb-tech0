package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"rest-api/src/database"
	"rest-api/src/handlers"
	"rest-api/src/nats/pub"
	"rest-api/src/nats/sub"
	"rest-api/src/service"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/nats-io/nats.go"
)

func main() {
	log.Println("Starting an application")

	// getting environment variables
	err := godotenv.Load()

	if err != nil {
		log.Printf("Some error occured. Err: %s", err)
	}

	db := database.MustRun()

	repo := database.NewOrderRepo(db)

	service := service.New(repo)

	handlers := handlers.New(service)

	repo.UpdateCacheOnStartUp()

	go func() {
		if err := http.ListenAndServe(":3002", handlers.InitHandlers()); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			pub.Publish()
		}
	}()

	var natsConn *nats.Conn
	go func() {
		natsConn = sub.Subscribe(repo)
	}()

	// graceful shutdown
	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	log.Println("stopping application")

	if err = db.Client.Close(); err != nil {
		log.Println("Error when closing database connection: ", err)
	}

	if err = natsConn.Drain(); err != nil {
		log.Println("Error when closing nats connection: ", err)
	}

	log.Println("application stopped")
}
