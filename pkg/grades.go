package pkg

import "strconv"

// Calculating student grades
func StudentStatus(absence, p1, p2, p3 string) string {
	absenceInt, _ := strconv.ParseFloat(absence, 64)
	p1Int, _ := strconv.Atoi(p1)
	p2Int, _ := strconv.Atoi(p2)
	p3Int, _ := strconv.Atoi(p3)
	absences := int((absenceInt / 60) * 100)
	average := ((p1Int / 10) + (p2Int / 10) + (p3Int / 10)) / 3
	if absences >= 25 {
		return "Reprovado por Falta"
	}

	if average >= 7 {
		return "Aprovado"
	} else if average < 7 && average >= 5 {
		return "Exame Final"
	} else {
		return "Reprovado por Nota"
	}
}

// Verify if student is in finals
func FinalExames(res, p1, p2, p3 string) string {
	p1Int, _ := strconv.ParseFloat(p1, 64)
	p2Int, _ := strconv.ParseFloat(p2, 64)
	p3Int, _ := strconv.ParseFloat(p3, 64)

	average := int(((p1Int / 10) + (p2Int / 10) + (p3Int / 10)) / 3)

	if res != "Exame Final" {
		return "0"
	}

	return strconv.Itoa(((average + 5) / 2))
}
