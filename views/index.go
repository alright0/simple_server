package views

import (
	"context"
	"fmt"
	"main/db"
	"net/http"

	"github.com/jackc/pgx/v5"
)

func IndexView(w http.ResponseWriter, r *http.Request) {

	con, err := db.Connect()
	if err != nil {
		http.Error(w, "error while connecting to db", http.StatusInternalServerError)
		return
	}
	defer con.Close()

	rows, err := con.Query(context.Background(), "select * from main")
	if err != nil {
		http.Error(w, "error while querying db", http.StatusInternalServerError)
	}

	type Us struct {
		ID   int    `db:"id"`
		Name string `db:"name"`
	}

	res, err := pgx.CollectRows(rows, pgx.RowToStructByName[Us])
	if err != nil {
		http.Error(w, "error while getting rows from db", http.StatusInternalServerError)
		return
	}
	fmt.Println(res)

	r.Header.Set("Content-Type", "application/json")
	fmt.Fprintf(w, string("Connect is successful"))

}
