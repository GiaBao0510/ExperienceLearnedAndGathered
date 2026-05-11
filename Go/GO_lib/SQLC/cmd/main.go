package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sqlc/internal/database"

	_ "github.com/lib/pq"
)

func main() {

	connecStr := "host=localhost port=5432 user=admin password=admin123 dbname=test sslmode=disable"

	pgDB, err := sql.Open("postgres", connecStr)
	if err != nil {
		panic(err)
	}
	defer pgDB.Close()

	// Excution
	dao := database.New(pgDB)
	ctx := context.Background()

	// Insert user
	err = dao.CreateUsers(ctx, database.CreateUsersParams{
		Uuid:        "1212",
		UserName:    "John",
		Email:       "john@example.com",
		PhoneNumber: "1234567890",
		Password:    "123456",
	})
	if err != nil {
		log.Fatal(err)
	}

	//get list
	userList, err := dao.GetAllUserss(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(userList)
}
