package main

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

func main() {
    // Replace with your actual database credentials
    // connStr := "postgres://postgres:postgres@localhost:5432/postgres"
    connStr := "user=postgres password=postgres dbname=postgres host=localhost port=5432 sslmode=disable"


    db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    // Test the connection
    err = db.Ping()
    if err != nil {
        panic(err)

    }

    fmt.Println("Connected to PostgreSQL!")

    // Example CRUD Operations:
    _, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name students, age INTEGER)")
    if err != nil {
        panic(err)
    }
    fmt.Println("table created")
        // Create a new record
        _, err = db.Exec("INSERT INTO students (column1, column2) VALUES ($1, $2)", "value1", "value2")
        if err != nil {
                panic(err)
        }

        // Read a record
        var result string
        err = db.QueryRow("SELECT column1 FROM students WHERE id = $1", 1).Scan(&result)
        if err != nil {
                panic(err)
        }
        fmt.Println(result)

        // Update a record
        _, err = db.Exec("UPDATE students SET column1 = $1 WHERE id = $2", "new_value", 1)
        if err != nil {
                panic(err)
        }

        // Delete a record
        _, err = db.Exec("DELETE FROM students WHERE id = $1", 1)
        if err != nil {
                panic(err)
        }


}