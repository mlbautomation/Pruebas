package carpeta

import "fmt"

type alumno struct {
	nombre string
}

func Saludar() {

	ML := alumno{
		nombre: "Marlon Ly",
	}

	fmt.Println("hola:", ML)
}
