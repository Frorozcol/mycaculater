package mycaculater

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

func (calc) operate(entrada string, operador string) int {

	entradaLimpia := strings.Split(entrada, operador)
	operacion1 := parsear(entradaLimpia[0])
	operacion2 := parsear(entradaLimpia[1])

	switch operador {
	case "+":
		fmt.Println(operacion1 + operacion2)
		return operacion1 + operacion2
	case "-":
		fmt.Println(operacion1 - operacion2)
		return operacion1 - operacion2
	case "*":
		fmt.Println(operacion1 * operacion2)
		return operacion1 * operacion2
	case "/":
		fmt.Println(operacion1 / operacion2)
		return operacion1 / operacion2
	default:
		fmt.Println(operador, "No está soportado")
		return 0
	}

}

func parsear(valor string) int {
	operacion, _ := strconv.Atoi(valor)
	return operacion
}

func leerEntrada(option string) string {
	if option == "entrada" {
		fmt.Println("Ingresé operación")
	} else {
		fmt.Println("Ingrese operador (+,-,*,/)")
	}
	scnaner := bufio.NewScanner(os.Stdin)
	scnaner.Scan()
	return scnaner.Text()
}

func main() {
	entrada := leerEntrada("entrada")
	operador := leerEntrada("operador")
	c := calc{}
	fmt.Println(c.operate(entrada, operador))

}
