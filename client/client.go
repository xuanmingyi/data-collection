package main

import (
	"context"

	"github.com/xuanmingyi/data-collection/client/ent"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	client, err := ent.Open("sqlite3", "file:a.db?_fk=1")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		panic(err)
	}
}
