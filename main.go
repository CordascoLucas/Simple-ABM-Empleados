package main

import(
	"net/http"
	"log"
	"text/template"
)

type Empleado struct {
	Id int
	Nombre string
	Correo string
}

var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/create", Create)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/delete", Delete)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/update", Update)

	log.Println("Servidor corriendo en el puerto 8000")
	http.ListenAndServe(":8000", nil)
}




