package main

import (
	"database/sql"
	"fmt"
    /* The '_' underscore before the SQLite driver import is used
       to include the package without directly using its functions. 
       This is necessary to ensure that the driver registers itself 
       with the database/sql package. */
	_ "github.com/mattn/go-sqlite3"
)

func main () {

    db, err := sql.Open("sqlite3", "./prophecy.db")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("open successfully!")
    }
    defer db.Close()

    rows, err := db.Query("SELECT * FROM students_import WHERE id = '1'")
    if err != nil {
        fmt.Println(err)
        return
    }
    var id, name, house, head string

    Columns, err := rows.Columns()
    if err != nil {
        fmt.Println("err at rows.Columns:", err)
    }
    fmt.Println(Columns)
    for rows.Next() {
        err = rows.Scan(&id, &name, &house, &head)
        if err != nil {
            fmt.Println("err at rows.Scan:", err)
        }
        fmt.Println(id, name, house, head)
    }

}

