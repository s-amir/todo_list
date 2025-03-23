package test

func DayOfWeek(i int8) string {
	switch i {
	case 1:
		return "Monday"
	case 2:
		return "Tuesday"
	case 3:
		return "Wednesday"
	case 4:
		return "Thursday"
	case 5:
		return "Friday"
	case 6:
		return "Saturday"
	case 7:
		return "Sunday"
	default:
		return ""
	}
}

func Square(i float64) float64 {
	return i * i
}
