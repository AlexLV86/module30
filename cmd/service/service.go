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
	db, err := storage.New("postgres://postgres:" + pwd + "@192.168.1.62:/tasks")
	if err != nil {
		log.Fatal(err)
	}
	tasks, err := db.Tasks(0, -1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tasks)
	err = db.DeleteTask(1)
	if err != nil {
		log.Fatal(err)
	}
	// users, err := db.Users(0)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(users)
	// taskid := 2
	// err = db.DeleteTask(taskid)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
