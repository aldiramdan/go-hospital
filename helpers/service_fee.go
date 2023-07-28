package helpers

import "github.com/aldiramdan/hospital/databases/db/models"

func ServiceFeeDoctor(data models.Diseases, fee float64) float64 {
	sum := 0.0
	for _, v := range data {
		sum += v.Medicine.Price
	}

	result := sum + fee + 15000.0

	return result
}
