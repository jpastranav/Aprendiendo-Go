package main // Declaración del paquete principal

import (
	"fmt"
	learnGO "goApp/learningGO"
) // Importación del paquete fmt para imprimir en consola

var alias string = "JX123" // Declaración de variable global con var y tipo explícito

//alias := "JX123" // Declaración de variable global con := no es válida, se debe usar var

func Var_Arrays_Slices() {

	// var primerNombre string = "Juan"
	// var segundoNombre string = "Perez"
	// var edad int8 = 30
	// Declaración de variables con var y tipo explícito 

	primerNombre := "Juan"
	segundoNombre := "Pastrana"
	edad := 28

	fmt.Println(primerNombre)
	fmt.Println(segundoNombre)
	fmt.Println(edad)
	fmt.Println(alias)

	//var lista = []string{"Juan", "Pastrana", "28"} -- Slice de tamaño variable declarado con var
	//lista := [3]string{"Juan", "Pastrana", "28"} -- Arreglo de tamaño fijo

	lista := []string{"Juan", "Pastrana", "28"} // Slice de tamaño variable declarado con :=
	fmt.Println(lista)

	lista = append(lista, "Colombia") // Agregar un elemento al slice
	fmt.Println(lista)

	mostrarPaises := lista[1:3] // Slicing: obtener un sub-slice desde el índice 1 hasta el índice 3 (excluyendo el índice 3)
	fmt.Println(mostrarPaises)

	mostrarPaises2 := lista[:3] // Slicing: obtener un sub-slice desde el inicio hasta el índice 3 (excluyendo el índice 3)
	fmt.Println(mostrarPaises2)

	mapa := map[string]string{"nombre": "Juan", "apellido": "Pastrana"} // Declaración de un mapa con clave y valor de tipo string
	fmt.Println(mapa["nombre"])                                         // Acceso al valor del mapa mediante la clave "nombre"
	mapa["edad"] = "28"                                                 // Agregar un nuevo par clave-valor al mapa
	fmt.Println(mapa)
}

func Condicionales() {
	edad := 18

	if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	} else {
		fmt.Println("Eres menor de edad")
	}

	switch edad {
	case 18:
		fmt.Println("Tienes 18 años")
	case 19:
		fmt.Println("Tienes 19 años")
	default:
		fmt.Println("No tienes 18 ni 19 años")
	}
}

func Bucles() {
	for i := 0; i < 5; i++ { // Bucle for con inicialización, condición y post-incremento
		fmt.Println(i)
	}

	j := 0 // Bucle while simulado con for y condición // While, no existe en Go, se simula con for y condición
	for j < 5 {
		fmt.Println(j)
		j++
	}

	k := 0 // Bucle infinito simulado con for y condición
	for {
		if k >= 5 {
			break
		}
		fmt.Println(k)
		k++
	}
}

func Funciones() {
	// Función que recibe parámetros y retorna un valor
	sumar := func(a int, b int) int {
		return a + b
	}

	resultado := sumar(2, 3)
	fmt.Println(resultado)

	// Función que no recibe parámetros y no retorna valor
	imprimir := func() {
		fmt.Println("Hola desde una función anónima")
	}

	imprimir()
}

func FuncionAnonima() {
	// Función anónima que se ejecuta inmediatamente
	func() {
		fmt.Println("Hola desde una función anónima que se ejecuta inmediatamente")
	}()

	// Función anónima asignada a una variable
	anonima := func() {
		fmt.Println("Hola desde una función anónima asignada a una variable")
	}

	anonima()

	//funcion closure: función anónima que captura variables del entorno
	contador := 0
	incrementar := func() int {
		contador++
		return contador
	}

	fmt.Println(incrementar()) // 1
	fmt.Println(incrementar()) // 2
	fmt.Println(incrementar()) // 3

	//funcion con currying: función que retorna otra función
	sumar := func(a int) func(b int) int {
		return func(b int) int {
			return a + b
		}
	}

	sumar5 := sumar(5)
	fmt.Println(sumar5(3)) // 8

	//funcion con parámetros variádicos: función que recibe un número variable de argumentos
	sumarVariadico := func(numeros ...int) int {
		suma := 0
		for _, numero := range numeros {
			suma += numero
		}
		return suma
	}

	fmt.Println(sumarVariadico(1, 2, 3, 4, 5)) // 15
}

func Calculadora() {

	// Declaración de un mapa de funciones anónimas para realizar operaciones matemáticas
	var operaciones = map[string]func(int, int) int{
		"suma": func(a int, b int) int {
			return a + b
		},
		"resta": func(a int, b int) int {
			return a - b
		},
		"multiplicacion": func(a int, b int) int {
			return a * b
		},
		"division": func(a int, b int) int {
			if b == 0 {
				fmt.Println("Error: división por cero")
				return 0
			}
			return a / b
		},
	}

	var operacion string
	var a, b int

	// Solicitar al usuario que ingrese la operación y los números
	fmt.Println("Ingrese la operación (suma, resta, multiplicacion, division):")
	fmt.Scanln(&operacion) // &operacion es un puntero a la variable operacion, lo que permite modificar su valor dentro de la función Scanln.
	fmt.Println("Ingrese el primer número:")
	fmt.Scanln(&a)
	fmt.Println("Ingrese el segundo número:")
	fmt.Scanln(&b)

	// Verificar si la operación ingresada es válida y ejecutar la función correspondiente
	if op, ok := operaciones[operacion]; ok {
		resultado := op(a, b)
		fmt.Printf("El resultado de la %s es: %d\n", operacion, resultado)
	} else {
		fmt.Println("Operación no válida")
	}
}

func FuncionImportada(){
	fmt.Println(learnGO.FuncionSuma(5,4))
}

/*func main() {
	var_arrays_slices()
	condicionales()
	bucles()
	funciones()
	funcionAnonima()
	calculadora()
}*/