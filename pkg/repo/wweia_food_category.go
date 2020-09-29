package repo

//go:generate reform

// WweiaFoodCategory represents a row in wweia_food_category table.
//reform:wweia_food_category
type WweiaFoodCategory struct {
	WweiaFoodCategoryCode        *string `reform:"wweia_food_category_code"`
	WweiaFoodCategoryDescription *string `reform:"wweia_food_category_description"`
}
