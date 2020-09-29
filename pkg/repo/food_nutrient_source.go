package repo

//go:generate reform

// FoodNutrientSource represents a row in food_nutrient_source table.
//reform:food_nutrient_source
type FoodNutrientSource struct {
	ID          *string `reform:"id"`
	Code        *string `reform:"code"`
	Description *string `reform:"description"`
}
