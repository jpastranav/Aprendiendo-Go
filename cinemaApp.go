package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var nombresPersona string
var apellidosPersona string
var edadInput int

var recaudoInfantil int
var recaudoAdultoMayor int
var recaudoGeneral int

var cantidadPersonas int

func runApp(scanner *bufio.Scanner) {

	for {
		fmt.Println("Cantidad de personas a ingresar: ")
		scanner.Scan()
		cantidadString := strings.TrimSpace(scanner.Text())

		cantidad, err := strconv.Atoi(cantidadString)

		if err != nil || cantidad < 0 {
			fmt.Println("Ingrese una cantidad valida de personas. No puede ser 0 o menor.")
			continue
		}
		cantidadPersonas = cantidad
		break
	}

	fmt.Println("Ingrese nombres: ")
	scanner.Scan()
	nombreInput := strings.TrimSpace(scanner.Text())

	if nombreInput == "" {
		fmt.Println("No se puede guardar nombre vacío. Se guardará 'N.N'")
		nombresPersona = "N.N"
		fmt.Println(nombresPersona)
	} else {
		fmt.Println("Nombres registrados: ", nombreInput)
		nombresPersona = nombreInput
	}

	fmt.Println("Ingrese apellidos:")
	scanner.Scan()
	apellidoInput := strings.TrimSpace(scanner.Text())

	if apellidoInput == "" {
		fmt.Println("No se puede guardar nombre vacío. Se guardará 'N.N'")
		apellidosPersona = "N.N"
		fmt.Println(apellidosPersona)
	} else {
		fmt.Println("Apellidos registrados: ", apellidoInput)
		apellidosPersona = apellidoInput
	}

	if nombreInput == "N.N" && apellidoInput == "N.N" {
		fmt.Println("Usuario registrado: ", nombreInput, " ", apellidoInput)
	} else {
		fmt.Print("Usuario registrado: ", nombresPersona, " ", apellidosPersona)
	}

	for {
		fmt.Println("Ingrese edad: ")
		scanner.Scan()
		edadString := strings.TrimSpace(scanner.Text())

		edad, err := strconv.Atoi(edadString)

		if err != nil || edad < 1 {
			fmt.Println("Edad no válida, intente en un rango mayor que 0")
			continue
		}
		edadInput = edad
		break
	}

	if edadInput < 12 {
		fmt.Println("Menor de 12 años: INFANTIL")
		recaudoInfantil += 4
	} else if edadInput >= 60 {
		fmt.Println("Mayor de 60 años: ADULTO MAYOR")
		recaudoAdultoMayor += 5
	} else {
		fmt.Println("Entre 12 - 59 años: GENERAL")
		recaudoGeneral += 7
	}

	recaudoTotal := recaudoInfantil + recaudoGeneral + recaudoAdultoMayor

	fmt.Println("El recaudo total de INFANTIL es: $", recaudoInfantil)
	fmt.Println("El recaudo total de ADULTO MAYOR es: $", recaudoAdultoMayor)
	fmt.Println("El recaudo total de GENERAL es: $", recaudoGeneral)

	fmt.Println("El recaudo TOTAL es: ", recaudoTotal)

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	runApp(scanner)
}
