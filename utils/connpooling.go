package utils

import (
	"connection_pooling/types"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func PoolExample(wg *sync.WaitGroup, pool *pgxpool.Pool, i int) {
	fmt.Println("added 1 more Job , i = ", i)
	time.Sleep(5 * time.Second)
	var u types.User
	query := "select * from connections;"
	rows, err := pool.Query(context.Background(), query)
	if err != nil {
		fmt.Println(fmt.Errorf("i = %d , error is %s ", i, err.Error()))
	}
	defer rows.Close()

	for rows.Next() {
		time.Sleep(5 * time.Second)
		rows.Scan(&u.Id, &u.Name, &u.Age)
		fmt.Printf("i = %d, name = %s, id = %d, age = %d ", i, u.Name, u.Id, u.Age)
		fmt.Println()
		fmt.Println()
	}

	wg.Done()

}
