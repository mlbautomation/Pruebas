package main

import (
	"fmt"

	carpeta1 "github.com/mlbautomation/Pruebas/carpeta1"
	carpeta2 "github.com/mlbautomation/Pruebas/carpeta2"
	"github.com/mlbautomation/Pruebas/errores"
)

func main() {
	persona1 := UserModel{
		nombre: "Ernesto",
		edad:   68,
	}
	/*
		var storage1 Storage

		usuario1 := New(storage1)

		err := usuario1.Create(&persona1)
		if err != nil {
			fmt.Println("usuario1.Create(&persona1):", err)
			return
		}

		fmt.Println("ejecutado correctamente", usuario1)
	*/

	fmt.Println("Hola mundo!", persona1)

	// mira en import...
	carpeta1.Saludar()
	carpeta2.Saludar()

	min := 3
	max := 100
	minConn := "10"
	maxConn := "9"
	verify(minConn, maxConn, &min, &max)
	fmt.Println(min, max)

	// veamos algo de errores personalizados
	m1 := errores.NewMessage1("debe ser un error!")
	m2 := errores.NewMessage2("ojo!")
	fmt.Printf("tipo: %T, valor: %v\n", m1, m1)
	fmt.Printf("tipo: %T, valor: %v\n", m2, m2)
	fmt.Println(errores.VerificacionError(m1))
	//errores.VerificacionError(m2)
}
