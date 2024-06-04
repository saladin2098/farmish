package config

func CalcWaterForPoultryPerDay(day float64) float64 {
	baseDosage := 5.3
	multiplier := 1 + (day/4)*0.4
	dosage := baseDosage * multiplier * 1.4
	return dosage
}

func CalcFoodForPoultryPerDay(day float64) float64 {
	baseDosage := 0.05
	multiplier := 1 + (day/7)*0.3
	dosage := baseDosage * multiplier
	return dosage
}

func CalcWaterPerDay(day float64, weight float64) float64 {
	baseDosage := 0.05
	multiplier := 1 + (day/100)*0.1
	dosage := baseDosage * multiplier * weight
	return dosage
}

func CalcFoodPerDay(day float64, weight float64) float64 {
	baseDosage := 0.1
	multiplier := 1 + (day/200)*0.1
	dosage := baseDosage * multiplier * weight
	return dosage
}

func CalcConsumptionForAnimals(day float64, weight float64) (float64, float64) {
	waterConsumption := CalcWaterPerDay(day, weight)
	foodConsumption := CalcFoodPerDay(day, weight)
	return waterConsumption, foodConsumption
}
