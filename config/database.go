package config

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "log"
)

//membuat function yang mereturn koneksi database
func GetConnection() *sql.DB {
  db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/data_mhs")
  if err != nil {
    log.Fatal(err)
  }
  
  return db
}