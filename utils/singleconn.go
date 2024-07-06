package utils

import (
	"context"
	"fmt"
	"sync"
	"time"

	"connection_pooling/types"

	"github.com/jackc/pgx/v5"
)

func CreateConnAndExecute(wg *sync.WaitGroup, conn *pgx.Conn, i int) {
	var u types.User
	query := "select * from connections;"
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		fmt.Println(fmt.Errorf("i = %d , error is %s ", i, err.Error()))
	}
	defer rows.Close()
	time.Sleep(3 * time.Second)
	for rows.Next() {
		rows.Scan(&u.Id, &u.Name, &u.Age)
		fmt.Printf("i = %d, name = %s, id = %d, age = %d ", i, u.Name, u.Id, u.Age)
	}
	wg.Done()
}
