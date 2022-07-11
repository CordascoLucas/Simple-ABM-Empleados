package main

import (
  "os"
  "github.com/joho/godotenv"
  "database/sql"
  _"github.com/go-sql-driver/mysql"
)

func ConexionBD()(conexion *sql.DB){
  err_dot := godotenv.Load("variables.env")

  if err_dot!=nil{
    panic(err_dot.Error())
  }

	Driver := os.Getenv("DRIVER")
	Usuario := os.Getenv("DBUSER")
	Contrasenia := os.Getenv("PASSWORD")
  Nombre := os.Getenv("DBNAME")
  IP := os.Getenv("IP")

	conexion, err := sql.Open(Driver, Usuario+":"+Contrasenia+"@tcp("+IP+")/"+Nombre)
	if err!=nil{
		panic(err.Error())
	}
	return conexion
}
