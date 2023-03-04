package tariff

func HrtoMinutes(hr float64) float64 {
	return hr * 60
}

func HrtoDays(hr float64) float64 {
	return hr / 24
}

func MintoDays(min float64) float64 {
	return HrtoDays(MintoHr(min))
}

func DaytoMinutes(day float64) float64 {
	return HrtoMinutes(day * 24)
}

func MintoHr(min float64) float64 {
	return min / 60
}
