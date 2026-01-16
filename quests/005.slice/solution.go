package slice

func ProcessScores(scores []int) []int {
	processed := []int{}
	for _, score := range scores {
		if score < 0 || score > 100 {
			continue
		} else if score < 40 {
			score = 40
		}

		processed = append(processed, score)
	}

	bonusPoints := len(processed) > 5
	for i, score := range processed {
		if bonusPoints {
			score += 5
		}
		if score > 100 {
			score = 100
		}

		processed[i] = score
	}
	return processed
}
