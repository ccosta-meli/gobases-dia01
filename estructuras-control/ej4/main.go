/*
Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. Según el siguiente mapa, ayudá a imprimir la edad de Benjamin.

var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

Por otro lado también es necesario:
	- Saber cuántos de sus empleados son mayores de 21 años.
	- Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
	- Eliminar a Pedro del mapa.
*/

package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	for nombre, edad := range employees {
		if edad > 21 {
			fmt.Println(nombre, "-> Mayor a 21 años, tiene:", edad)
		}
	}

	fmt.Println("Agregamos a Federico de 25 años y eliminamos a Pedro")

	employees["Federico"] = 25

	delete(employees, "Pedro")

	for nombre, edad := range employees {
		fmt.Println("Nombre", nombre, "-> Edad:", edad)
	}
}
