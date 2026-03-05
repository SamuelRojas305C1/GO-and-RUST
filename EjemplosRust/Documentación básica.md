
## **Hola mundo
```
fn main() {

    println!("¡Hola, mundo desde Rust!")

}
```
En Rust, el punto de entrada es siempre la función `main`.

fn: Palabra clave para declarar una función.
    
println!: Fíjate en el signo de exclamación !. Esto indica que no es una función normal, sino una **macro**. Las macros en Rust se expanden en código más complejo durante la compilación para permitir cosas que las funciones normales no pueden (como un número variable de argumentos).
    
Compilación: A diferencia de lenguajes interpretados, Rust se compila a código máquina puro.


## **Variables y Mutabilidad
```
fn main() {

    let lenguaje = Rust;

  

    let mut version = 1;

    version = 2;

  

    println!("Estamos aprendiendo {} version {}", lenguaje, version);

}
```
Este es uno de los conceptos más importantes de Rust: la **inmutabilidad por defecto**.

let: Declara una variable. Por defecto, una vez que le asignas un valor, no puedes cambiarlo. Esto evita errores de concurrencia.
    
mut: Si necesitas que una variable cambie, debes ser explícito y usar let mut.
    
Inferencia de tipos: Rust es de tipado estático, pero casi siempre adivina qué tipo de dato estás usando (como `i32` para enteros o &str para texto) sin que tengas que escribirlo.


## **Ciclos
```
fn main() {

  

    for i in 1..6{

        println!("Iteracion numero: {}",i);

    }

}
```
Rust prefiere un enfoque más moderno y seguro para los ciclos, evitando los errores de "fuera de índice".

Rangos (1..6): Representa una secuencia de números. El primer número es inclusivo y el último es exclusivo (va del 1 al 5). Para incluir el 6, usarías 1..=6.
    
Iteradores: El ciclo for en Rust no funciona incrementando un contador manualmente como en C; funciona pidiendo el siguiente elemento a un "iterador", lo que lo hace mucho más eficiente y seguro.


## **Funciones y Expresiones
```
fn sumar(a: i32, b: i32) -> i32 {

    a + b

}

fn main() {

    let resultado = sumar(10, 20);

    println!("La suma es: {}", resultado);

}
```
Rust distingue claramente entre **sentencias** (instrucciones que terminan en ;) y **expresiones** (bloques de código que devuelven un valor).

Firma de función Los tipos de parámetros (a: i32) y el retorno -> i32 son obligatorios en las funciones.

Retorno implícito: En Rust, no necesitas escribir `return` si el valor que quieres devolver es la última expresión de la función y **no tiene punto y coma**. Esto hace el código más limpio.

 a + b es una expresión (devuelve valor). a + b; es una sentencia (no devuelve nada).


## **Condicionales
```
fn main() {

    let edad = 18;

  

    if edad >= 18 {

        println!("Eres mayor de edad");

    } else {

        println!("Eres menor de edad");

    }

}
```
El if en Rust funciona de manera lógica, pero tiene un superpoder: puede usarse para asignar valores.

Sin paréntesis: Al igual que en Go, no los necesita, lo que mejora la lectura.

Asignación por if: Puedes hacer algo como: let resultado = if x > 5 { "grande" } else { "pequeño" };

Tipos estrictos: La condición debe ser obligatoriamente un booleano. Rust no considera que el número `1` sea "true" o `0` sea "false".

