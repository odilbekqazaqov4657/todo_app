package main

import (
	"fmt"
	"log"

	"github.com/odilbekqazaqov4657/todo_app/pkg/db"
	"honnef.co/go/tools/config"
)

func main() {
	cfg := config.LLoad()
	con, err := db.NewPgx(cfg)

	if err != nil {
		return
	}

	log.Println("Successfully connected to db !")

	fmt.Println(con)
}
