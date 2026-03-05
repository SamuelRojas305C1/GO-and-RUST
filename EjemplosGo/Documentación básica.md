
## **Hola mundo
```
package main

import fmt

func main() {

    fmt.Println("¡Hola, mundo desde Go!")

}
```
En Go, el concepto de **paquete** es la unidad básica de organización.

 package main: Indica al compilador de Go que este archivo debe compilarse como un **ejecutable** y no como una librería compartida.
    
import fmt: Importa el paquete "Format", que contiene funciones para la entrada/salida de texto (I/O).
    
func main(): Es el punto de entrada único. Cuando ejecutas el programa, el sistema busca esta función para empezar.
    
 En Go, la llave de apertura `{` **debe** estar en la misma línea que la declaración de la función o el bucle, de lo contrario, el compilador arrojará un error.


## **Asignación de variable
```
Package main

import fmt

  

func main() {

    var lenguaje string = "Go"

  

    version := 1

  

    fmt.Printf("Estamos aprendiendo %s version %d\n", lenguaje, version)

}
```
Go ofrece seguridad de tipos (no puedes sumar un texto con un número), pero es flexible en cómo los declaras.

var: Forma explícita. Útil cuando declaras una variable sin asignarle un valor inmediato (se inicializa en su "zero value", como `0` para números o `""` para strings).
    
Operador de inferencia (`:=`): Solo se puede usar dentro de funciones. Go deduce el tipo de dato basándose en el valor de la derecha.
    
Tipado estático: Una vez que una variable es definida como `int`, no puede cambiar a `string` durante la ejecución.


## **Ciclos
```
Package main

import fmt

  

func main() {

    for i := 1; i<=5; i++ {

        fmt.Println("Iteracion numero:", i)

    }

}
```
El for es la única palabra clave para iteraciones en Go, lo que reduce la complejidad del lenguaje.

| **Estructura**                  | **Uso**                              |
| ------------------------------- | ------------------------------------ |
| for i := 0; i < 10; i++         | Estilo clásico (C-style).            |
| for condicion { ... }           | Funciona como un while.              |
| for { ... }                     | Ciclo infinito (se rompe con break). |
| for index, value := range lista | Para recorrer arreglos o mapas.      |


## **Funciones
```
Package

import fmt

  

func sumar(a int, b int) int {

    return a + b

}

func main() {

    resultado := sumar(10, 20)

    fmt.Prinln("La suma es:", resultado)

}
```
Las funciones son ciudadanos de primera clase en Go (pueden pasarse como variables).

Parámetros: Se definen como nombre tipo. Si varios parámetros seguidos son del mismo tipo, puedes agruparlos: func sumar(a, b int).
    
Retorno: El tipo de dato que devuelve la función se coloca justo antes de la llave de apertura {.
    
Múltiples retornos: Go permite devolver varios valores (muy usado para devolver un resultado y un error simultáneamente): func division(a, b int) (int, error).


## **Condicionales
```
package main

import fmt

  

func main() {

    edad := 18

  

    if edad >= 18{

        fmt.Println("Eres mayor de edad")

    } else {

        fmt.Println("Eres menor de edad")

    }

}
```
El flujo de control en Go busca la legibilidad por encima de todo.

Sin paréntesis: A diferencia de C o Java, no se usan () para la condición, lo que limpia visualmente el código.
    
Scope (Alcance): Go permite una instrucción corta antes de la condición:
```
if valor := calcular(); valor > 10 {
    fmt.Println(valor)
}
```

Booleanos estrictos: La condición debe ser estrictamente un valor booleano (true o false). No puedes usar números (como if 1) como se hace en otros lenguajes.