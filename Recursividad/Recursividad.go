package main;

import(
	"fmt";
	"bufio";
	"os";
	"strconv";
	"strings";
);

func main() {
	scanner := bufio.NewReader(os.Stdin);

	fmt.Printf("Ingresa un número para ver su factorial: ");

	var linea string = "";
	var numero float64 = 0;

	linea, err := scanner.ReadString('\n');
	// Excepcion de lectura
	if(err != nil) {
		fmt.Fprintln(os.Stderr, "Error de lectura: ", err);
		os.Exit(2);
	}

	// Sustitucion de retorno de carro y salto de linea
	linea = strings.Replace(linea, "\r\n", "", -1);

	// Conversion de cadena
	numero, err2 := strconv.ParseFloat(linea, 64);
	
	// Excepcion de conversion
	if(err2 != nil) {
		fmt.Fprintln(os.Stderr, "Error de conversion: ", err2);
		os.Exit(2);
	}

	// Impresion de resultado
	fmt.Printf("El factorial de %s es: %v", linea, getFactorial(numero));
}

func getFactorial(numero float64) float64 {
	if(numero == 1) {
		return 1;
	
	} else {
		return numero * getFactorial(numero - 1);
	}
}

/**
 * Por Fernando Florez
*/