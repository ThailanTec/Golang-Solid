package main

import "reflect"

// Modelo errado de se fazer esse calculo.

type funcCLT struct {
}
type Estagiario struct{}

func FolhaDePag[TipoFuncionario funcCLT | Estagiario](funcionario TipoFuncionario) float64 {

	if reflect.TypeOf(funcCLT{}) == reflect.TypeOf(funcionario) {
		// calcular salario + beneficios
		return 0.0
	} else {
		//Calcula bolsa auxilio estágiario
		return 0.0
	}

}

func main() {

}
