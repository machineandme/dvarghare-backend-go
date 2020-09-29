package repo

//go:generate reform

// FoodProteinConversionFactor represents a row in food_protein_conversion_factor table.
//reform:food_protein_conversion_factor
type FoodProteinConversionFactor struct {
	FoodNutrientConversionFactorID *string `reform:"food_nutrient_conversion_factor_id"`
	Value                          *string `reform:"value"`
}
