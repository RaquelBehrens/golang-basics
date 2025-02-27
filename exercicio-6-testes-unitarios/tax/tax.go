package tax

func Tax(salario float64) (resultado float64) {
	resultado = 0
	if salario >= 50000 {
		resultado += (salario * 0.17)
	}
	if salario >= 150000 {
		resultado += (salario * 0.10)
	}
	return
}
