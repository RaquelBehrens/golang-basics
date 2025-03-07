package grades

func Average(notas ...float64) float64 {
	var sum float64 = 0
	for _, nota := range notas {
		sum += nota
	}
	return sum / float64(len(notas))
}
