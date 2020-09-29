package repo

//go:generate reform

// FoodAttributeType represents a row in food_attribute_type table.
//reform:food_attribute_type
type FoodAttributeType struct {
	ID          *string `reform:"id"`
	Name        *string `reform:"name"`
	Description *string `reform:"description"`
}
