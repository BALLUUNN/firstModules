package utils

const DayOfWeek int = 7

func CountOfDay(week int) int {
	return week * DayOfWeek
}

func DaysToWeeks(days int) float64 {
	return float64(days) / float64(DayOfWeek)
}
