package main

import "net/http"

func Index(w http.ResponseWriter, r *http.Request){
	conexionEstablecida := ConexionBD()
	registros, err := conexionEstablecida.Query("SELECT * FROM empleados")

	if err!=nil {
		panic(err.Error())
	}

	empleado := Empleado{}
	arregloEmpleado := []Empleado{}

	for registros.Next(){
		var id int
		var nombre, correo string
		err = registros.Scan(&id, &nombre, &correo)

		if err!=nil{
			panic(err.Error())
		}

		empleado.Id = id
		empleado.Nombre = nombre
		empleado.Correo = correo

		arregloEmpleado=append(arregloEmpleado, empleado)
	}

	plantillas.ExecuteTemplate(w, "index", arregloEmpleado)
}

func Create(w http.ResponseWriter, r *http.Request){
	plantillas.ExecuteTemplate(w, "create", nil)
}

func Insert(w http.ResponseWriter, r *http.Request){
	if r.Method=="POST"{
		nombre := r.FormValue("name")
		correo := r.FormValue("email")

		conexionEstablecida := ConexionBD()
		insertarRegistros, err := conexionEstablecida.Prepare("INSERT INTO empleados(nombre, correo) VALUES(?, ?)")

		if err!=nil {
			panic(err.Error())
		}

		insertarRegistros.Exec(nombre, correo)

		http.Redirect(w,r,"/",301)
	}
}

func Delete(w http.ResponseWriter, r *http.Request){
	idEmpleado := r.URL.Query().Get("id")

	conexionEstablecida := ConexionBD()
	borrarRegistro, err := conexionEstablecida.Prepare("DELETE FROM empleados WHERE id = ?")

	if err!=nil {
		panic(err.Error())
	}

	borrarRegistro.Exec(idEmpleado)

	http.Redirect(w,r,"/",301)
}

func Edit(w http.ResponseWriter, r *http.Request){
	idEmpleado := r.URL.Query().Get("id")

	conexionEstablecida := ConexionBD()
	registro, err := conexionEstablecida.Query("SELECT * FROM empleados WHERE id=?", idEmpleado)

	if err!=nil {
		panic(err.Error())
	}

	empleado := Empleado{}

	for registro.Next(){
		var id int
		var nombre, correo string
		err = registro.Scan(&id, &nombre, &correo)

		if err!=nil{
			panic(err.Error())
		}

		empleado.Id = id
		empleado.Nombre = nombre
		empleado.Correo = correo

	}

	plantillas.ExecuteTemplate(w, "edit", empleado)
}

func Update(w http.ResponseWriter, r *http.Request){
	if r.Method=="POST"{
		id := r.FormValue("id")
		nombre := r.FormValue("name")
		correo := r.FormValue("email")

		conexionEstablecida := ConexionBD()
		modificarRegistro, err := conexionEstablecida.Prepare("UPDATE empleados SET nombre=?, correo=? WHERE id=?")

		if err!=nil {
			panic(err.Error())
		}

		modificarRegistro.Exec(nombre, correo, id)

		http.Redirect(w,r,"/",301)
	}
}
