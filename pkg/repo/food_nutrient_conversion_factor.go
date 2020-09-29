package repo

//go:generate reform

// FoodNutrientConversionFactor represents a row in food_nutrient_conversion_factor table.
//reform:food_nutrient_conversion_factor
type FoodNutrientConversionFactor struct {
	ID    *string `reform:"id"`
	FdcID *string `reform:"fdc_id"`
}
