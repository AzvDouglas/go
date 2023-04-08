package tax

func CalculateTax(amount float64) float64 {
	if amount >= 70 {
		return 10
	} 
	return 5
}