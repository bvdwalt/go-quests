package conditions

func ClassifyRequest(age int, hasID bool, balance float64, isVIP bool) string {
	if age < 0 || balance < 0 {
		return "INVALID"
	} else if age < 18 || !hasID {
		return "REJECTED"
	} else if isVIP && balance >= 10000 {
		return "VIP_ACCESS"
	} else if balance >= 1000 {
		return "STANDARD_ACCESS"
	} else {
		return "LIMITED_ACCESS"
	}
}

func EvaluateGrade(score int) string {
	if score < 0 || score > 100 {
		return "INVALID"
	} else if score >= 90 {
		return "A"
	} else if score >= 80 {
		return "B"
	} else if score >= 70 {
		return "C"
	} else if score >= 60 {
		return "D"
	} else {
		return "F"
	}
}
