package repo

//go:generate reform

// FoodCategory represents a row in food_category table.
//reform:food_category
type FoodCategory struct {
	ID          *string `reform:"id"`
	Code        *string `reform:"code"`
	Description *string `reform:"description"`
}
