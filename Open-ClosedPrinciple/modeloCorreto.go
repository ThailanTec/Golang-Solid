package main

type FuncPJ struct{}

type FuncCLT struct{}

type FuncEstagiario struct{}

type Funcionario interface {
	CalcularSalario() float64
}

func (f *FuncPJ) CalcularSalario() float64 {
	// Calcular Salario
	return 8.000
}
func (f *funcCLT) CalcularSalario() float64 {
	// Calcular Salario
	return 4.000
}

func (f *FuncEstagiario) CalcularSalario() float64 {
	// Calcular Salario
	return 800.00
}

func FolhaDePagamento(funcionario Funcionario) float64 {
	return funcionario.CalcularSalario()
}
