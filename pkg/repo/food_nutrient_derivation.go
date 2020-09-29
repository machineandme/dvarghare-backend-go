package repo

//go:generate reform

// FoodNutrientDerivation represents a row in food_nutrient_derivation table.
//reform:food_nutrient_derivation
type FoodNutrientDerivation struct {
	ID          *string `reform:"id"`
	Code        *string `reform:"code"`
	Description *string `reform:"description"`
	SourceID    *string `reform:"source_id"`
}
