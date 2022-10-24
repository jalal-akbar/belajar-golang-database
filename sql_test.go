package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Insert Data
func TestExecDatabase(t *testing.T) {

	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	script := "INSERT INTO customers(id, name)VALUES('1', 'Jalal');"

	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert Data to Database")
}

// Query ke Database
func TestQuery(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "SELECT * FROM customers;"

	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	// Iterasi
	for rows.Next() {
		var id, name string
		// Membaca data rows
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("id :", id)
		fmt.Println("name :", name)
	}
	defer rows.Close()
}

// Tipe Data Column
// Erorr Parsing time.Time
// Error NULL VALUE
func TestQueryComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	sqlScript := "SELECT * FROM customers;"
	rows, err := db.QueryContext(context.Background(), sqlScript)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var (
			id, name   string
			email      sql.NullString // tipe data untuk NULL value
			balance    int64
			rating     float64
			created_at time.Time
			birth_date sql.NullTime // tipe data untuk null value
			married    bool
		)
		err := rows.Scan(&id, &name, &email, &balance, &rating, &created_at, &birth_date, &married)
		if err != nil {
			panic(err)
		}
		// mengecek null atau tidak
		fmt.Println("id :", id)
		fmt.Println("name :", name)
		// mengecek null atau tidak
		if email.Valid {
			fmt.Println("Email: ", email)
		}
		fmt.Println("balance:", balance)
		fmt.Println("rating:", rating)
		// mengecek null atau tidak
		if birth_date.Valid {
			fmt.Println("Birth Date :", birth_date)
		}
		fmt.Println("married :", married)
		fmt.Println("created_at :", created_at)
	}
	rows.Close()
}

// SQL Injection
func TestSQLInjection(t *testing.T) {
	// sql dengan injection
	username := "jalal'; #" // celah dari query injection
	// hardcode = SELECT username FROM users WHERE username = 'jalal'; # AND passwords = 'SALAH' LIMIT 1;
	// setelah 'jalal'; akan di hiraukan karna #
	passwords := "SALAH" // Sukses Login meskipun passwordnya salah

	var sqlQueryInjection = "SELECT username FROM users WHERE username = '" + username + "' AND passwords ='" + passwords + "' LIMIT 1;"
	db := GetConnection()
	defer db.Close()
	rows, err := db.QueryContext(context.Background(), sqlQueryInjection)
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login")

	} else {
		fmt.Println("Gagal Login")
	}
	rows.Close()
}

// SQL dengan paramater
// Solusi dari SQL Injection
func TestInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	var username = "jalal"
	var passwords = "SALAH"
	var scriptSQL = "SELECT username FROM users WHERE username = ? AND passwords = ? LIMIT 1;" // gunakan ? untuk inject yg aman
	rows, err := db.QueryContext(context.Background(), scriptSQL, username, passwords)
	if err != nil {
		panic(err)
	}
	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login")
	} else {
		fmt.Println("Gagal Login")
	}
	defer rows.Close()
}
