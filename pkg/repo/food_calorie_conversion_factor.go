package repo

//go:generate reform

// FoodCalorieConversionFactor represents a row in food_calorie_conversion_factor table.
//reform:food_calorie_conversion_factor
type FoodCalorieConversionFactor struct {
	FoodNutrientConversionFactorID *string `reform:"food_nutrient_conversion_factor_id"`
	ProteinValue                   *string `reform:"protein_value"`
	FatValue                       *string `reform:"fat_value"`
	CarbohydrateValue              *string `reform:"carbohydrate_value"`
}
