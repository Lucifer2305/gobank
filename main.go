package main

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", store)

	//server := NewAPIServer(":3000")
	//server.Run()
	//fmt.Println("Yeah Buddy")

}

/*const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "password"
	dbname   = "root"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}*/
