package main

import (
	"fmt"
	"hl4-user_service/handler"
	"hl4-user_service/pkg/server"
	"hl4-user_service/pkg/store"
	"hl4-user_service/repository"
	"hl4-user_service/service"
	"os"
)

func main() {
	dsn := os.Getenv("DSN")

	fmt.Println(dsn)
	insertExampleData := true
	db, err := store.New(dsn, insertExampleData)
	if err != nil {
		return
	}
	defer db.Client.Close()

	rep := repository.New(db.Client)
	serv := service.New(rep)
	handlers := handler.New(serv)

	httpServ, err := server.New(handlers.HTTP, "8080")
	if err != nil {
		return
	}
	fmt.Println("start to run...")
	if err = httpServ.Http.ListenAndServe(); err != nil {
		return
	}
	fmt.Println("runnin...")

	quit := make(chan os.Signal, 1)

	<-quit
}
