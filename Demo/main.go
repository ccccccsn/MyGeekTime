package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"strconv"
)

func StudentDAO(uid int) (bool, error) {
	var name string
	db, err := sql.Open("sqlite3", "db/school.db")
	if err != nil {
		return false, errors.Wrapf(err, "Open db failed")
	}
	Query := "SELECT name FROM student WHERE id = " + strconv.Itoa(uid)
	err = db.QueryRow(Query).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows{
			return false, errors.Wrapf(err, "error no rows. sql: %s", Query)
		}else{
			return false, errors.Wrapf(err, "QueryRow failed. sql: %s", Query)
		}
	}

	defer db.Close()
	return true, nil
}

func main()  {
	//IsFound, err := StudentDAO(1)
	IsFound, err := StudentDAO(7)
	if err != nil{
		fmt.Printf("fail: %+v\n", err)
	}
	fmt.Println("IsFound: ", IsFound)
}