/*
02 - Clima.
Una empresa de meteorología quiere una aplicación donde pueda tener la temperatura, humedad y presión atmosférica de distintos lugares.
	1. Declará 3 variables especificando el tipo de dato, como valor deben tener la temperatura, humedad y presión de donde te encuentres.
	2. Imprimí los valores de las variables en consola.
	3. ¿Qué tipo de dato le asignarías a las variables?
*/

package main

import "fmt"

func main() {
	var temp float64 = 11
	var humedad int = 58
	var presion int = 1_022

	fmt.Println("Temperatura:", temp)
	fmt.Println("Humedad:", humedad)
	fmt.Println("Presion:", presion)
}
