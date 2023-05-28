package main

import (
    "database/sql"
    "fmt"
    "log"
)

func gay() {
    // Create a new database connection.
    db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/mydb")
    if err != nil {
        log.Fatal(err)
    }

    // Close the database connection when we're done.
    defer db.Close()

    // Create a new statement.
    stmt, err := db.Prepare("SELECT * FROM users")
    if err != nil {
        log.Fatal(err)
    }

    // Execute the statement.
    rows, err := stmt.Query()
    if err != nil {
        log.Fatal(err)
    }

    // Iterate over the rows and print the data.
    for rows.Next() {
        var (
            id int
            name string
        )

        err := rows.Scan(&id, &name)
        if err != nil {
            log.Fatal(err)
        }

        fmt.Printf("ID: %d, Name: %s\n", id, name)
    }
}