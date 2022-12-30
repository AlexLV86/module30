package main

import (
	"fmt"
	"log"
	"os"

	"module30/module307/pkg/storage"
)

func main() {
	pwd := os.Getenv("dbpass")
	if pwd == "" {
		os.Exit(1)
	}
	fmt.Println("Password: ", pwd)
	db, err := storage.New("postgres://postgres:1234@192.168.1.62:/tasks")
	if err != nil {
		log.Fatal(err)
	}
	tasks, err := db.Tasks(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tasks)
	users, err := db.Users(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(users)
}
