/*
Préstamo

Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos. Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar. Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un año de antigüedad en su trabajo. Dentro de los préstamos que otorga no les cobrará interés a los que posean un sueldo superior a $100.000.

Es necesario realizar una aplicación que reciba  estas variables y que imprima un mensaje de acuerdo a cada caso.

Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.
*/

package main

import "fmt"

func main() {
	var edad int = 23
	var sueldo float64 = 100_000
	var empleado bool = true
	var antiguedad int = 1

	if !empleado {
		fmt.Println("El usuario no esta empleado")
	} else if edad <= 22 {
		fmt.Println("El usuario es menor de 22 años")
	} else if antiguedad < 1 {
		fmt.Println("El usuario no cumple con la antiguedad requerida")
	} else if sueldo >= 100_000 {
		fmt.Println("El usuario puede acceder al prestamo, no se le cobrara intereses.")
	} else {
		fmt.Println("El usuario puede acceder al prestamos, se le cobrara intereses")
	}
}
