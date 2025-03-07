package alunos

import (
	"fmt"
	"time"
)

type Student struct {
	Name          string
	Surname       string
	DNI           int
	AdmissionDate time.Time
}

func (s Student) Detalhamento() {
	ad := s.AdmissionDate
	fmt.Printf("Name: %s, Surname: %s, Admission Date: %d/%.2d/%d, DNI: %d\n", s.Name, s.Surname, ad.Day(), ad.Month(), ad.Year(), s.DNI)
}
