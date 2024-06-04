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

func CalcWaterForAnimalsPerDay(day float64, weight float64) float64 {
	baseDosage := 0.035
	multiplier := 1 + (day/100)*0.1
	dosage := baseDosage * multiplier * weight
	return dosage
}

func CalcFoodForAnimalsPerDay(day float64, weight float64) float64 {
	baseDosage := 0.040
	multiplier := 1 + (day/200)*0.1
	dosage := baseDosage * multiplier * weight
	return dosage
}

func CalcConsumptionForPoultry(day float64) (float64, float64) {
	waterConsumption := CalcWaterForPoultryPerDay(day)
	foodConsumption := CalcFoodForPoultryPerDay(day)
	return waterConsumption, foodConsumption
}

func CalcConsumptionForAnimals(day float64, weight float64) (float64, float64) {
	waterConsumption := CalcWaterForAnimalsPerDay(day, weight)
	foodConsumption := CalcFoodForAnimalsPerDay(day, weight)
	return waterConsumption, foodConsumption
}
