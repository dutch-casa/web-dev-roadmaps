package main

type Student struct {
	Name    string
	Email   string
	Scores  []int
	Tuition int // in cents
}

func averageScore(scores []int) float64 {
	total := 0
	for _, s := range scores {
		total += s
	}
	return float64(total) / float64(len(scores))
}
