package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"quirky-store-go/internal/chaos"
	"quirky-store-go/internal/repository"
	"quirky-store-go/internal/service"
)

func main() {

	db, err := sql.Open(
		"mysql",
		"username:password@tcp(localhost:3306)/quirky_store",
	)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewStoreRepository(db)
	ch := &chaos.DefaultChaos{}
	store := service.NewQuirkyStore(repo, ch)

	_ = store.Put("cat", "meow")
	_ = store.Put("dog", "woof")
	_ = store.Put("eagle", "Fly high")
	_ = store.Put("snake", "Crawling")

	v, ok, _ := store.Get("cat")

	fmt.Println(v, ok)

	rows, _ := store.Dump()
	fmt.Println(rows)
}
