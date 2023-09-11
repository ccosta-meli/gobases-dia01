/*
A que mes corresponde.

Realizar una aplicación que reciba  una variable con el número del mes.
	1. Según el número, imprimir el mes que corresponda en texto.
	2. ¿Se te ocurre que se puede resolver de distintas maneras? ¿Cuál elegirías y por qué?
Ej: 7, Julio.
Nota: Validar que el número del mes, sea correcto.
*/

package main

import "fmt"

func main() {
	var nroMes int = 7

	switch nroMes {
	case 1:
		fmt.Printf("%d, Enero \n", nroMes)
	case 2:
		fmt.Printf("%d, Febrero \n", nroMes)
	case 3:
		fmt.Printf("%d, Marzo \n", nroMes)
	case 4:
		fmt.Printf("%d, Abril \n", nroMes)
	case 5:
		fmt.Printf("%d, Mayo \n", nroMes)
	case 6:
		fmt.Printf("%d, Junio \n", nroMes)
	case 7:
		fmt.Printf("%d, Julio \n", nroMes)
	case 8:
		fmt.Printf("%d, Agosto \n", nroMes)
	case 9:
		fmt.Printf("%d, Septiembre \n", nroMes)
	case 10:
		fmt.Printf("%d, Octubre \n", nroMes)
	case 11:
		fmt.Printf("%d, Noviembre \n", nroMes)
	case 12:
		fmt.Printf("%d, Diciembre \n", nroMes)

	}

}
