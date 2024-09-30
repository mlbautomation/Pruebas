package carpeta

import "fmt"

type alumno struct {
	nombre string
}

func Saludar() {

	JL := alumno{
		nombre: "Johao Ly",
	}

	fmt.Println("hola:", JL)
}
