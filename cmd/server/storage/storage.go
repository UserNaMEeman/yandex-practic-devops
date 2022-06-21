package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DataStore struct {
	Name   string
	Type   string
	ValueF float64
	ValueC int64
}

func (data *DataStore) SaveData() {
	db, err := sql.Open("mysql", "root:root@/Metrics")
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%T\n", db)
	defer db.Close()
	val := DataStore{}
	// fmt.Println("do")
	rows, errdb := db.Query("select * from metrics where name = ?", data.Name)
	// rows, errdb := db.Query("select * from Metrics.gauge where name = ?", "Alloc")
	if errdb != nil {
		panic(errdb)
	}
	defer rows.Close()
	// fmt.Println("ok")
	if rows.Next() {
		err := rows.Scan(&val.Name, &val.Type, &val.ValueF, &val.ValueC)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		// fmt.Println(val.Type)
		switch val.Type {
		case "gauge":
			// fmt.Println("gauge")
			_, errdb := db.Exec("update metrics set valueGauge = ? where name = ?", data.ValueF, data.Name)
			fmt.Println(data.Name, data.ValueF)
			if errdb != nil {
				panic(errdb)
			}
		case "counter":
			// fmt.Println("counter")
			newVal := data.ValueC + val.ValueC
			_, errdb := db.Exec("update metrics set valueCounter = ? where name = ?", newVal, data.Name)
			if errdb != nil {
				panic(errdb)
			}
		}
	} else {
		_, errdb := db.Exec("insert into metrics (name, type, valueGauge, valueCounter) values (?, ?, ?, ?)", data.Name, data.Type, data.ValueF, data.ValueC)
		if errdb != nil {
			panic(errdb)
		}
	}
}

// func (data *Counter)SaveDataD(){
// 	var randV int64 = 7				//Change string - add storage method
// 	data.Value = data.Value + randV
// 	// fmt.Println(data.PollCount)
// }
