package main

import (
	"context"
	"fmt"
	"os"
	"sync"

	"connection_pooling/utils"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func pgxPoolExample() {
	connString := "postgres://postgres:postgres@localhost:5432/pooling?application_name=conn_pooling_example"
	// pgxConfig := pgxpool.Config{
	// 	MaxConns: 5,
	// 	MinConns: 2,
	// }
	var wg sync.WaitGroup
	pgxConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		fmt.Println("error is ", err)
	}
	pgxConfig.MaxConns = 10
	pgxConfig.MinConns = 5
	dbpool, err := pgxpool.NewWithConfig(context.Background(), pgxConfig)
	if err != nil {
		fmt.Println("error is ", err)
	}
	defer dbpool.Close()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		// time.Sleep(1 * time.Second)
		go utils.PoolExample(&wg, dbpool, i)
	}
	wg.Wait()
}
func pgxSingleConnExample() {
	connString := "postgres://postgres:postgres@localhost:5432/pooling?application_name=conn_pooling_example"
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		fmt.Println("error is ", err)
		os.Exit(1)
	}
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go utils.CreateConnAndExecute(&wg, conn, i)
	}
	wg.Wait()
}
func main() {
	//pgxSingleConnExample()
	pgxPoolExample()
}
