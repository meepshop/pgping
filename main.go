package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/cenkalti/backoff"

	_ "github.com/lib/pq"
)

func main() {
	connectionString := os.Args[1]
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	backoff.Retry(func() error {
		err := db.Ping()
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Print("OK")
		return nil
	}, backoff.NewExponentialBackOff())
}
